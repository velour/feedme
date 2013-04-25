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

type Article struct {
	Title, Link, OriginTitle string
	DescriptionData          []byte
	When                     time.Time
}

func (a Article) Description() template.HTML {
	return template.HTML(a.DescriptionData)
}

type FeedInfo struct {
	Title     string
	Url       string
	Articles  []Article
	Refs      int
	LastFetch time.Time
}

func (f FeedInfo) Len() int {
	return len(f.Articles)
}

func (f FeedInfo) Less(i, j int) bool {
	return f.Articles[i].When.After(f.Articles[j].When)
}

func (f FeedInfo) Swap(i, j int) {
	f.Articles[i], f.Articles[j] = f.Articles[j], f.Articles[i]
}

// UrlKey returns the key for the given URL.
func urlKey(c appengine.Context, url string) *datastore.Key {
	return datastore.NewKey(c, "Feed", url, 0, nil)
}

// getFeeds returns an aggregate feed.
func getFeeds(c appengine.Context, keys []*datastore.Key) (FeedInfo, error) {
	var as []Article
	for _, k := range keys {
		f, err := getFeedByKey(c, k)
		if err != nil {
			return FeedInfo{}, err
		}
		as = append(as, f.Articles...)
	}
	return FeedInfo{Title: "All Feeds", Articles: as}, nil
}

// GetFeedByKey loads a feed from the datastore with the given key.
// If the Feed has not been fetched from the source within MaxCacheTime
// then the Feed is read from it's source and added to the store.
func getFeedByKey(c appengine.Context, key *datastore.Key) (FeedInfo, error) {
	var f FeedInfo
	err := datastore.Get(c, key, &f)
	if err != nil {
		return FeedInfo{}, err
	}
	if time.Since(f.LastFetch) > maxCacheDuration {
		f, err = fetchAndStoreFeed(c, f.Url)
	}
	return f, err
}

// GetFeedByUrl loads a feed from the datastore with the given URL.
// If the Feed is not in the store or has not been fetched from the source
// within MaxCacheTime then the Feed is read from it's source and added
// to the store.
func getFeedByUrl(c appengine.Context, url string) (FeedInfo, error) {
	var f FeedInfo
	key := urlKey(c, url)

	err := datastore.Get(c, key, &f)
	if err != nil && err != datastore.ErrNoSuchEntity {
		return FeedInfo{}, err
	}
	if err == datastore.ErrNoSuchEntity || time.Since(f.LastFetch) > maxCacheDuration {
		f, err = fetchAndStoreFeed(c, url)
	}
	return f, err
}

func fetchAndStoreFeed(c appengine.Context, url string) (FeedInfo, error) {
	feed, err := fetchFeed(c, url)
	if err != nil {
		return FeedInfo{}, err
	}

	key := urlKey(c, url)

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

	return feed, err
}

// FetchFeed reads a feed from the given URL.
func fetchFeed(c appengine.Context, url string) (FeedInfo, error) {
	resp, err := urlfetch.Client(c).Get(url)
	if err != nil {
		return FeedInfo{}, err
	}
	defer resp.Body.Close()

	feed, err := rss.Get(resp.Body)
	if err != nil {
		return FeedInfo{}, err
	}
	as := make([]Article, len(feed.Items))
	for i, item := range feed.Items {
		as[i] = Article{
			Title:           item.Title,
			Link:            item.Link,
			OriginTitle:     feed.Title,
			DescriptionData: []byte(item.Description),
			When:            item.When,
		}
	}
	return FeedInfo{
		Title:     feed.Title,
		Url:       url,
		Articles:  as,
		LastFetch: time.Now(),
	}, nil
}
