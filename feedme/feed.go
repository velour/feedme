package feedme

import (
	"appengine"
	"appengine/datastore"
	"appengine/urlfetch"
	"html/template"
	"time"

	rss "github.com/zippoxer/RSS-Go"
)

// maxCacheDuration is the length of time to store a feed before refetching it
const maxCacheDuration = time.Minute * 30

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

// GetFeedByKey loads a feed from the datastore with the given key.
// If the Feed has not been fetched from the source within MaxCacheTime
// then the Feed is read from it's source and added to the store.
func getFeedInfoByKey(c appengine.Context, key *datastore.Key) (FeedInfo, error) {
	var f FeedInfo
	err := datastore.Get(c, key, &f)
	if err != nil {
		return FeedInfo{}, err
	}
	if time.Since(f.LastFetch) > maxCacheDuration {
		f, err = refreshFeed(c, f.Url)
	}
	return f, err
}

// GetFeedByUrl loads a feed from the datastore with the given URL.
// If the Feed is not in the store or has not been fetched from the source
// within MaxCacheTime then the Feed is read from it's source and added
// to the store.
func getFeedInfoByUrl(c appengine.Context, url string) (FeedInfo, error) {
	var f FeedInfo
	key := datastore.NewKey(c, "Feed", url, 0, nil)

	err := datastore.Get(c, key, &f)
	if err != nil && err != datastore.ErrNoSuchEntity {
		return FeedInfo{}, err
	}
	if err == datastore.ErrNoSuchEntity || time.Since(f.LastFetch) > maxCacheDuration {
		f, err = refreshFeed(c, url)
	}
	return f, err
}

// GetArticles gets all articles for a feed.
func getArticles(c appengine.Context, feedKeys ...*datastore.Key) (articles Articles, err error) {
	for _, f := range feedKeys {
		if _, err = datastore.NewQuery("Article").Ancestor(f).GetAll(c, &articles); err != nil {
			break
		}
	}
	return
}

// RefreshFeed fetches the feed from the remote source, stores it's info and articles in
// the datastore, removes old articles (those not retrieved on the latest fetch), and
// returns its FeedInfo.
func refreshFeed(c appengine.Context, url string) (FeedInfo, error) {
	feed, articles, err := fetchFeed(c, url)
	if err != nil {
		return FeedInfo{}, err
	}

	feedKey := datastore.NewKey(c, "Feed", url, 0, nil)

	// Add a new FeedInfo or update the LastFetch time.
	err = datastore.RunInTransaction(c, func(c appengine.Context) error {
		var f FeedInfo
		err := datastore.Get(c, feedKey, &f)
		if err == datastore.ErrNoSuchEntity {
			err = nil
			f = feed
		}
		if err != nil {
			return err
		}
		feed.Refs = f.Refs
		_, err = datastore.Put(c, feedKey, &feed)
		return err
	}, nil)
	if err != nil {
		return FeedInfo{}, err
	}

	err = updateArticles(c, feedKey, articles)

	return feed, err
}

// RmArticles removes the articles associated with a feed.
func rmArticles(c appengine.Context, feedKey *datastore.Key) error {
	q := datastore.NewQuery("Article").Ancestor(feedKey).KeysOnly()
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
	q := datastore.NewQuery("Article").Ancestor(feedKey).KeysOnly()
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
		k := datastore.NewKey(c, "Article", a.Link, 0, feedKey)
		id := k.StringID()
		if _, ok := stored[id]; ok {
			delete(stored, id)
			continue
		}
		if _, err := datastore.Put(c, k, a); err != nil {
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
