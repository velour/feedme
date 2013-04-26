package feedme

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"time"
)

var (
	templateFiles = []string{
		"tmplt/list.html",
		"tmplt/feed.html",
	}

	funcs = template.FuncMap{
		"dateTime":   func(t time.Time) string { return t.Format("2006-01-02 15:04:05") },
		"timeString": func(t time.Time) string { return t.Format("Mon Jan 2 15:04:05 MST 2006") },
	}

	templates = template.Must(template.New("t").Funcs(funcs).ParseFiles(templateFiles...))
)

const maxFeeds = 50

type UserInfo struct {
	Feeds []*datastore.Key
}

func init() {
	http.HandleFunc("/list", handleList)
	http.HandleFunc("/add", handleAdd)
	http.HandleFunc("/rm", handleRm)
	http.HandleFunc("/", handleRoot)
}

type root struct {
	User  UserInfo
	Feeds []userFeedInfo
}

// FeedInfo is the information about a feed in the user's feed list.
type userFeedInfo struct {
	Title      string
	LastFetch  time.Time
	EncodedKey string
}

func (f userFeedInfo) Fresh() bool {
	return time.Since(f.LastFetch) < maxCacheDuration
}

func handleList(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/list" || r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	c := appengine.NewContext(r)

	uinfo, err := userInfo(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var feeds []userFeedInfo
	for _, key := range uinfo.Feeds {
		var feed FeedInfo
		err := datastore.Get(c, key, &feed)
		if err == datastore.ErrNoSuchEntity {
			continue
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		feeds = append(feeds, userFeedInfo{
			Title:      feed.Title,
			LastFetch:  feed.LastFetch,
			EncodedKey: key.Encode(),
		})
	}

	if err := templates.ExecuteTemplate(w, "list.html", root{User: uinfo, Feeds: feeds}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	uinfo, err := userInfo(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(uinfo.Feeds) == 0 {
		http.Redirect(w, r, "/list", http.StatusFound)
		return
	}

	var feedPage = struct {
		Title    string
		Articles Articles
	}{}

	if r.URL.Path == "/" {
		feedPage.Title = "All Feeds"
		feedPage.Articles, err = getArticles(c, uinfo.Feeds...)
	} else {
		var key *datastore.Key
		if key, err = datastore.DecodeKey(r.URL.Path[1:]); err == nil {
			var f FeedInfo
			if f, err = getFeedInfoByKey(c, key); err == nil {
				feedPage.Title = f.Title
				feedPage.Articles, err = getArticles(c, key)
			}
		} else {
			err = fmt.Errorf("Failed to find feed")
		}
	}
	if err == datastore.ErrNoSuchEntity {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sort.Sort(feedPage.Articles)

	if err := templates.ExecuteTemplate(w, "feed.html", feedPage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	c := appengine.NewContext(r)

	// Check tha the feed is even valid, and put it in the datastore.
	url := r.FormValue("url")
	if _, err := getFeedInfoByUrl(c, url); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := addFeed(c, datastore.NewKey(c, "Feed", url, 0, nil)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/list", http.StatusFound)
}

func handleRm(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	c := appengine.NewContext(r)

	uinfo, err := userInfo(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, k := range uinfo.Feeds {
		if r.FormValue(k.Encode()) != "on" {
			continue
		}
		if err := rmFeed(c, k); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/list", http.StatusFound)
}

// AddFeed adds a feed to the user's feed list if it is not already there.
// The feed must already be in the datastore.
func addFeed(c appengine.Context, feedKey *datastore.Key) error {
	return datastore.RunInTransaction(c, func(c appengine.Context) error {
		u, err := userInfo(c)
		if err != nil {
			return err
		}

		if len(u.Feeds) >= maxFeeds {
			return fmt.Errorf("Too many feeds, max is %d", maxFeeds)
		}

		for _, k := range u.Feeds {
			if feedKey.Equal(k) {
				return nil
			}
		}

		var f FeedInfo
		if err := datastore.Get(c, feedKey, &f); err != nil {
			return err
		}

		f.Refs++
		if _, err := datastore.Put(c, feedKey, &f); err != nil {
			return err
		}

		u.Feeds = append(u.Feeds, feedKey)
		_, err = datastore.Put(c, userInfoKey(c), &u)
		return err
	}, &datastore.TransactionOptions{XG: true})
}

// RmFeed removes a feed from the user's feed list.
func rmFeed(c appengine.Context, feedKey *datastore.Key) error {
	return datastore.RunInTransaction(c, func(c appengine.Context) error {
		u, err := userInfo(c)
		if err != nil {
			return err
		}

		i := 0
		var k *datastore.Key
		for i, k = range u.Feeds {
			if feedKey.Equal(k) {
				break
			}
		}
		if i >= len(u.Feeds) {
			return nil
		}

		var f FeedInfo
		if err := datastore.Get(c, feedKey, &f); err != nil {
			return err
		}

		f.Refs--
		if f.Refs <= 0 {
			if err := rmArticles(c, feedKey); err != nil {
				return err
			}
			if err := datastore.Delete(c, feedKey); err != nil {
				return err
			}
		} else if _, err := datastore.Put(c, feedKey, &f); err != nil {
			return err
		}

		u.Feeds = append(u.Feeds[:i], u.Feeds[i+1:]...)
		_, err = datastore.Put(c, userInfoKey(c), &u)
		return err
	}, &datastore.TransactionOptions{XG: true})
}

// UserInfo returns the UserInfo for the currently logged in user.
// This function assumes that a user is loged in, otherwise it will panic.
func userInfo(c appengine.Context) (UserInfo, error) {
	var uinfo UserInfo
	err := datastore.Get(c, userInfoKey(c), &uinfo)
	if err != nil && err != datastore.ErrNoSuchEntity {
		return UserInfo{}, err
	}
	return uinfo, nil
}

// UserInfoKey returns the key for the current user's UserInfo.
// This function assumes that a user is loged in, otherwise it will panic.
func userInfoKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "User", user.Current(c).String(), 0, nil)
}
