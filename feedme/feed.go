package feedme

import (
	"appengine"
	"appengine/datastore"
	"appengine/urlfetch"
	"html/template"
	"sort"
	"time"

	rss "github.com/zippoxer/RSS-Go"
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

func makeFeedInfo(title, url string) FeedInfo {
	return FeedInfo{Title: title, Url: url}
}

// GetArticles returns all articles the feeds, refreshing them if necessary.
func getArticles(c appengine.Context, feeds ...FeedInfo) (articles Articles, err error) {
	for _, f := range feeds {
		key := datastore.NewKey(c, feedKind, f.Url, 0, nil)
		if time.Since(f.LastFetch) > maxCacheDuration {
			if err := refreshFeed(c, f); err != nil {
				break
			}
		}
		if _, err = datastore.NewQuery(articleKind).Ancestor(key).GetAll(c, &articles); err != nil {
			break
		}
	}
	return
}

// RefreshFeed fetches the feed from the remote source, stores it's info and articles in
// the datastore, removes old articles (those not retrieved on the latest fetch), and
// returns its FeedInfo.
func refreshFeed(c appengine.Context, f FeedInfo) error {
	feed, articles, err := fetchFeed(c, f.Url)
	if err != nil {
		return err
	}
	sort.Sort(articles)
	if len(articles) > maxNewArticles {
		articles = articles[:maxNewArticles]
	}

	key := datastore.NewKey(c, feedKind, f.Url, 0, nil)

	// Add a new FeedInfo or update the LastFetch time.
	err = datastore.RunInTransaction(c, func(c appengine.Context) error {
		var f FeedInfo
		err := datastore.Get(c, key, &f)
		if err == datastore.ErrNoSuchEntity {
			err = nil
			f = feed
		}
		if err != nil {
			return err
		}
		feed.Refs = f.Refs
		_, err = datastore.Put(c, key, &feed)
		return err
	}, nil)
	if err != nil {
		return err
	}

	return updateArticles(c, key, articles)
}

// RmArticles removes the articles associated with a feed.
func rmArticles(c appengine.Context, feedKey *datastore.Key) error {
	q := datastore.NewQuery(articleKind).Ancestor(feedKey).KeysOnly()
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

// UpdateArticles removes the old articles from the given feed and adds the new ones.
func updateArticles(c appengine.Context, feedKey *datastore.Key, articles []Article) error {
	q := datastore.NewQuery(articleKind).Ancestor(feedKey).KeysOnly()
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
		k := datastore.NewKey(c, articleKind, a.Link, 0, feedKey)
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

// FetchFeed reads a feed from the given URL.
func fetchFeed(c appengine.Context, url string) (FeedInfo, Articles, error) {
	var finfo FeedInfo
	resp, err := urlfetch.Client(c).Get(url)
	if err != nil {
		return finfo, nil, err
	}
	defer resp.Body.Close()

	feed, err := rss.Get(resp.Body)
	if err != nil {
		return finfo, nil, err
	}

	finfo.Title = feed.Title
	finfo.Url = url
	finfo.LastFetch = time.Now()

	as := make(Articles, len(feed.Items))
	for i, item := range feed.Items {
		as[i] = Article{
			Title:           item.Title,
			Link:            item.Link,
			OriginTitle:     feed.Title,
			DescriptionData: []byte(item.Description),
			When:            item.When,
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
	f, err := rss.Get(resp.Body)
	if err != nil {
		return "", err
	}
	return f.Title, err
}
