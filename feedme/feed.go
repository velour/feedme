package feedme

import (
	"appengine"
	"appengine/datastore"
	"appengine/urlfetch"
	"fmt"
	"github.com/velour/feedme/webfeed"
	"html/template"
	"sort"
	"time"
)

const (
	// MaxCacheDuration is the length of time to store a feed before refetching it
	maxCacheDuration = time.Minute * 30

	// MaxNewArticles is the maximum number of articles stored when fetching
	// new articles from a feed.
	maxNewArticles = 10

	articleKind = "Article"
	feedKind    = "Feed"
)

// An Article is a single article from a feed.
type Article struct {
	Title           string
	Link            string
	DescriptionData []byte
	When            time.Time
	// OriginTitle is the title of the feed from which this article originated.
	OriginTitle string
}

func (a Article) Description() template.HTML {
	return template.HTML(a.DescriptionData)
}

// Articles is a slice of Articles implementing sort.Interface.
type Articles []Article

func (as Articles) Len() int {
	return len(as)
}

func (as Articles) Less(i, j int) bool {
	return as[i].When.After(as[j].When)
}

func (as Articles) Swap(i, j int) {
	as[i], as[j] = as[j], as[i]
}

// FeedInfo is the information stored for each feed.
type FeedInfo struct {
	Title string
	Url   string

	// Refs is the number of users currently subscribed to the feed.
	Refs int

	// LastFetch is the last time the feed was fetched from the source.
	LastFetch time.Time
}

// MakeFeedInfo returns a FeedInfo with the given title and URL.
func makeFeedInfo(title, url string) FeedInfo {
	return FeedInfo{Title: title, Url: url}
}

// GetArticles returns all articles for a feed, refreshing it if necessary.
func (f FeedInfo) articles(c appengine.Context) (articles Articles, err error) {
	key := datastore.NewKey(c, feedKind, f.Url, 0, nil)
	if time.Since(f.LastFetch) > maxCacheDuration {
		if err = f.refresh(c); err != nil {
			return
		}
	}
	_, err = datastore.NewQuery(articleKind).Ancestor(key).GetAll(c, &articles)
	return
}

// RefreshFeed fetches the feed from the remote source, stores it's info and articles in
// the datastore, removes old articles (those not retrieved on the latest fetch), and
// returns its FeedInfo.
func (f FeedInfo) refresh(c appengine.Context) error {
	title, articles, err := f.readSource(c)
	if err != nil {
		return err
	}

	key := datastore.NewKey(c, feedKind, f.Url, 0, nil)
	err = datastore.RunInTransaction(c, func(c appengine.Context) error {
		var stored FeedInfo
		err := datastore.Get(c, key, &stored)
		if err != nil && err != datastore.ErrNoSuchEntity {
			return err
		}
		f.Title = title
		f.Refs = stored.Refs
		f.LastFetch = time.Now()
		_, err = datastore.Put(c, key, &f)
		return err
	}, nil)
	if err != nil {
		return err
	}

	return f.updateArticles(c, articles)
}

// ReadSource returns the feed title and articles read from the source.
func (f FeedInfo) readSource(c appengine.Context) (string, Articles, error) {
	feed, articles, err := fetchUrl(c, f.Url)
	if err != nil {
		return "", nil, err
	}
	sort.Sort(articles)
	if len(articles) > maxNewArticles {
		articles = articles[:maxNewArticles]
	}
	return feed.Title, articles, nil
}

func (f FeedInfo) updateArticles(c appengine.Context, articles Articles) error {
	key := datastore.NewKey(c, feedKind, f.Url, 0, nil)
	q := datastore.NewQuery(articleKind).Ancestor(key).KeysOnly()
	stored := make(map[string]*datastore.Key)
	for it := q.Run(c); ; {
		k, err := it.Next(nil)
		if err == datastore.Done {
			break
		} else if err != nil {
			return err
		}
		stored[k.StringID()] = k
	}

	for _, a := range articles {
		k := datastore.NewKey(c, articleKind, a.Link, 0, key)
		id := k.StringID()
		if _, ok := stored[id]; ok {
			delete(stored, id)
			continue
		}
		if _, err := datastore.Put(c, k, &a); err != nil {
			return err
		}
	}

	for _, k := range stored {
		if err := datastore.Delete(c, k); err != nil {
			return err
		}
	}
	return nil
}

// RmArticles removes the articles associated with a feed.
func (f FeedInfo) rmArticles(c appengine.Context) error {
	key := datastore.NewKey(c, feedKind, f.Url, 0, nil)
	q := datastore.NewQuery(articleKind).Ancestor(key).KeysOnly()
	for it := q.Run(c); ; {
		k, err := it.Next(nil)
		switch {
		case err == datastore.Done:
			return nil
		case err != nil:
			return err
		}
		if err := datastore.Delete(c, k); err != nil {
			return err
		}
	}
	return nil
}

// FetchUrl reads a feed from the given URL.
func fetchUrl(c appengine.Context, url string) (FeedInfo, Articles, error) {
	var finfo FeedInfo
	resp, err := urlfetch.Client(c).Get(url)
	if err != nil {
		return finfo, nil, err
	}
	defer resp.Body.Close()

	feed, err := webfeed.Read(resp.Body)
	if err != nil {
		err = fmt.Errorf("failed to fetch %s: %s\n", url, err.Error())
		return finfo, nil, err
	}

	finfo.Title = feed.Title
	finfo.Url = url
	finfo.LastFetch = time.Now()

	as := make(Articles, len(feed.Entries))
	for i, ent := range feed.Entries {
		content := ent.Content
		if len(ent.Content) == 0 {
			content = ent.Summary
		}
		as[i] = Article{
			Title:           ent.Title,
			Link:            ent.Link,
			OriginTitle:     feed.Title,
			DescriptionData: content,
			When:            ent.When,
		}
	}

	return finfo, as, nil
}

// CheckUrl returns the feed title and nil if the URL is a valid feed, otherwise it returns an error.
func checkUrl(c appengine.Context, url string) (string, error) {
	resp, err := urlfetch.Client(c).Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	f, err := webfeed.Read(resp.Body)
	if err != nil {
		return "", err
	}
	return f.Title, err
}
