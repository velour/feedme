package feedme

import (
	"appengine"
	"appengine/datastore"
	"appengine/memcache"
	"appengine/urlfetch"
	"errors"
	"github.com/velour/feedme/webfeed"
	"html/template"
	"net/http"
	"sort"
	"time"
)

const (
	// MaxCacheDuration is the length of time to store a feed before refetching it
	maxCacheDuration = 25 * time.Minute

	// MaxNewArticles is the maximum number of articles stored when fetching
	// new articles from a feed.
	maxNewArticles = 10

	// McacheKeyMax is the maximum key length for memcache.
	mcacheKeyMax = 250

	articleKind = "Article"
	feedKind    = "Feed"
)

var (
	// ErrKeyTooBig is returned if the key is too big for the memcache.
	ErrKeyTooBig = errors.New("Memcache key is too big")
)

// An Article is a single article from a feed.
type Article struct {
	Title           string    `datastore:",noindex"`
	Link            string    `datastore:",noindex"`
	DescriptionData []byte    `datastore:",noindex"`
	When            time.Time `datastore:",noindex"`
	WhenAdded       time.Time
	// OriginTitle is the title of the feed from which this article originated.
	OriginTitle string `datastore:",noindex"`
}

func (a Article) Description() template.HTML {
	return template.HTML(a.DescriptionData)
}

// StringID returns a unique string that can be used to identify this
// article in a datastore.Key.
func (a Article) StringID() string {
	return a.Title + a.When.Format("06-01-02;15:04:05")
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
	// The URL from which to fetch the Atom or RSS.
	Url string `datastore:",noindex"`

	Title string `datastore:",noindex"`
	Link  string `datastore:",noindex"`

	// Refs is the number of users currently subscribed to the feed.
	Refs int `datastore:",noindex"`

	// LastFetch is the last time the feed was fetched from the source.
	LastFetch time.Time `datastore:",noindex"`
}

// EnsureFresh refreshes the feed only if it is stale.
func (f *FeedInfo) ensureFresh(c appengine.Context) error {
	if time.Since(f.LastFetch) > maxCacheDuration {
		return f.refresh(c)
	}
	return nil
}

// RefreshFeed fetches and updates a feed from the remote source,
// stores it's info and articles in the datastore, and removes old articles
// (those not retrieved on the latest fetch).
func (f *FeedInfo) refresh(c appengine.Context) error {
	fnew, articles, fetchErr := f.readSource(c)

	err := datastore.RunInTransaction(c, func(c appengine.Context) error {
		var stored FeedInfo
		key := datastore.NewKey(c, feedKind, f.Url, 0, nil)
		err := fixMissingFieldError(datastore.Get(c, key, &stored))
		if err != nil && err != datastore.ErrNoSuchEntity {
			return err
		}
		if fetchErr != nil {
			*f = stored
			f.LastFetch = time.Now()
		} else {
			*f = fnew
		}
		f.Refs = stored.Refs
		_, err = datastore.Put(c, key, f)
		return err
	}, nil)
	if fetchErr != nil {
		return fetchErr
	}
	if err != nil {
		return err
	}

	return f.updateArticles(c, articles)
}

// GetArticles returns all articles for a feed, refreshing it if necessary.
func (f FeedInfo) articlesSince(c appengine.Context, t time.Time) (articles Articles, err error) {
	key := datastore.NewKey(c, feedKind, f.Url, 0, nil)
	if t.IsZero() {
		_, err = datastore.NewQuery(articleKind).Ancestor(key).GetAll(c, &articles)
	} else {
		_, err = datastore.NewQuery(articleKind).Ancestor(key).Filter("WhenAdded >=", t).GetAll(c, &articles)
	}
	return
}

// ReadSource returns the feed title and articles read from the source.
func (f FeedInfo) readSource(c appengine.Context) (FeedInfo, Articles, error) {
	feed, articles, err := fetchUrl(c, f.Url)
	if err != nil {
		return FeedInfo{}, nil, err
	}
	sort.Sort(articles)
	if len(articles) > maxNewArticles {
		articles = articles[:maxNewArticles]
	}
	return feed, articles, nil
}

func (f FeedInfo) updateArticles(c appengine.Context, articles Articles) error {
	feedKey := datastore.NewKey(c, feedKind, f.Url, 0, nil)
	articleKeys, err := f.articleKeys(c)
	if err != nil {
		return err
	}

	stored := make(map[string]*datastore.Key)
	for _, k := range articleKeys {
		stored[k.StringID()] = k
	}

	var nadded, nkept int
	debug := ""
	newKeys := make([]*datastore.Key, len(articles))
	for i, a := range articles {
		k := datastore.NewKey(c, articleKind, a.StringID(), 0, feedKey)
		newKeys[i] = k

		id := k.StringID()
		if _, ok := stored[id]; ok {
			debug += "Keeping " + id + "\n"
			nkept++
			delete(stored, id)
			continue
		}
		debug += "Adding " + id + "\n"
		nadded++
		a.WhenAdded = time.Now()
		if _, err := datastore.Put(c, k, &a); err != nil {
			return err
		}
	}

	ndeleted := len(stored)
	for _, k := range stored {
		debug += "Deleting " + k.StringID() + "\n"
		if err := datastore.Delete(c, k); err != nil {
			return err
		}
	}

	if nadded > 0 || ndeleted > 0 {
		c.Debugf("%s, added %d, kept %d, deleted %d\n%s", f.Url, nadded, nkept, ndeleted, debug)
	}

	item := memcache.Item{Key: articlesMcacheKey(feedKey), Object: newKeys}
	memcache.Gob.Set(c, &item)

	return nil
}

func (f FeedInfo) articleKeys(c appengine.Context) ([]*datastore.Key, error) {
	feedKey := datastore.NewKey(c, feedKind, f.Url, 0, nil)
	mcacheKey := articlesMcacheKey(feedKey)

	var keys []*datastore.Key

	if _, err := memcache.Gob.Get(c, mcacheKey, &keys); err == nil {
		return keys, nil
	}

	q := datastore.NewQuery(articleKind).Ancestor(feedKey).KeysOnly()
	for it := q.Run(c); ; {
		k, err := it.Next(nil)
		if err == datastore.Done {
			break
		} else if err != nil {
			return nil, err
		}
		keys = append(keys, k)
	}

	item := memcache.Item{Key: mcacheKey, Object: keys}
	memcache.Gob.Set(c, &item)

	return keys, nil
}

func articlesMcacheKey(feedKey *datastore.Key) string {
	return "articles:" + feedKey.StringID()
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
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return finfo, nil, errors.New("HTTP error: " + http.StatusText(resp.StatusCode))
	}
	defer resp.Body.Close()

	feed, err := webfeed.Read(resp.Body)
	if err != nil {
		if _, ok := err.(webfeed.ErrBadTime); ok {
			c.Infof("Ignoring bad time format for feed %s: %s", url, err.Error())
			err = nil
		} else {
			err = errors.New("failed to fetch " + url + ": " + err.Error())
			return finfo, nil, err
		}
	}

	finfo.Url = url
	finfo.Title = feed.Title
	if finfo.Title == "" {
		finfo.Title = url
	}
	finfo.Link = feed.Link
	finfo.LastFetch = time.Now()

	as := make(Articles, len(feed.Entries))
	for i, ent := range feed.Entries {
		content := ent.Content
		if len(ent.Content) == 0 {
			content = ent.Summary
		}
		title := ent.Title
		if title == "" && len(ent.Summary) > 0 {
			n := len(ent.Summary)
			if n > 20 {
				n = 20
			}
			title = string(ent.Summary[:n]) + "…"
		}
		if title == "" {
			title = ent.Link
		}
		as[i] = Article{
			Title:           title,
			Link:            ent.Link,
			OriginTitle:     feed.Title,
			DescriptionData: content,
			When:            ent.When,
		}
	}

	return finfo, as, nil
}

// CheckUrl returns information about a feed and nil if the URL is a
// valid feed, otherwise it returns an error.
func checkUrl(c appengine.Context, url string) (FeedInfo, error) {
	resp, err := urlfetch.Client(c).Get(url)
	if err != nil {
		return FeedInfo{}, err
	}
	defer resp.Body.Close()
	f, err := webfeed.Read(resp.Body)
	if err != nil {
		if _, ok := err.(webfeed.ErrBadTime); ok {
			c.Infof("Ignoring bad time format for feed %s: %s", url, err.Error())
			err = nil
		} else {
			return FeedInfo{}, err
		}
	}
	return FeedInfo{Url: url, Title: f.Title, Link: f.Link}, err
}

func getFeed(c appengine.Context, k *datastore.Key) (FeedInfo, error) {
	f, err := cacheGetFeed(c, k)
	if err == nil {
		return f, nil
	}

	if err := fixMissingFieldError(datastore.Get(c, k, &f)); err != nil {
		return FeedInfo{}, err
	}

	return f, cacheSetFeed(c, k, f)
}

func cacheGetFeed(c appengine.Context, k *datastore.Key) (FeedInfo, error) {
	id := k.StringID()
	if len(id) > mcacheKeyMax {
		return FeedInfo{}, ErrKeyTooBig
	}
	var f FeedInfo
	_, err := memcache.Gob.Get(c, id, &f)
	return f, err
}

func cacheSetFeed(c appengine.Context, k *datastore.Key, f FeedInfo) error {
	id := k.StringID()
	if len(id) > mcacheKeyMax { // silently don't set the value
		c.Infof("Not caching feed with a big key: %s", id)
		return nil
	}
	return memcache.Gob.Set(c, &memcache.Item{Key: k.StringID(), Object: f})
}

// Update updates a feed in a transaction by calling u with the stored
// version of the feed, and putting the feed returned by u.
func (f *FeedInfo) update(c appengine.Context, u func(stored FeedInfo) FeedInfo) error {
	k := datastore.NewKey(c, feedKind, f.Url, 0, nil)
	err := datastore.RunInTransaction(c, func(c appengine.Context) error {
		var stored FeedInfo
		err := fixMissingFieldError(datastore.Get(c, k, &stored))
		if err != nil && err != datastore.ErrNoSuchEntity {
			return err
		}
		*f = u(stored)
		_, err = datastore.Put(c, k, f)
		return err
	}, nil)
	if err != nil {
		return err
	}
	return cacheSetFeed(c, k, *f)
}

// Changes ErrFieldMismatch errors for the ArticleKeys field into nil.
func fixMissingFieldError(err error) error {
	if err == nil {
		return nil
	}

	switch e := err.(type) {
	case appengine.MultiError:
		var multi appengine.MultiError
		for _, err := range e {
			if err := fixMissingFieldError(err); err != nil {
				multi = append(multi, err)
			}
		}
		if len(multi) == 0 {
			return nil
		}
		return multi

	case *datastore.ErrFieldMismatch:
		if e.FieldName == "ArticleKeys" {
			return nil
		}
		return err
	}

	return err
}
