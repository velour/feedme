package feedme

import (
	"appengine"
	"appengine/datastore"
	"errors"
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

	uinfo, err := getUserInfo(c)
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

	uinfo, err := getUserInfo(c)
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
			err = errors.New("Failed to find feed")
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

	if err := subscribe(c, datastore.NewKey(c, "Feed", url, 0, nil)); err != nil {
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

	uinfo, err := getUserInfo(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, k := range uinfo.Feeds {
		if r.FormValue(k.Encode()) != "on" {
			continue
		}
		if err := unsubscribe(c, k); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/list", http.StatusFound)
}
