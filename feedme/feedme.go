package feedme

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strings"
	"time"
)

var (
	templateFiles = []string{
		"tmplt/list.html",
		"tmplt/feed.html",
	}

	funcs = template.FuncMap{
		"dateTime": func(t time.Time) string { return t.Format("2006-01-02 15:04:05") },
	}

	templates = template.Must(template.New("t").Funcs(funcs).ParseFiles(templateFiles...))
)

const (
	latestDuration = 18 * time.Hour
)

func init() {
	http.HandleFunc("/list", handleList)
	http.HandleFunc("/add", handleAdd)
	http.HandleFunc("/addopml", handleOpml)
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

// userFeedInfos is a type for sorting the infos.
type userFeedInfos []userFeedInfo

func (u userFeedInfos) Len() int {
	return len(u)
}

func (u userFeedInfos) Less(i, j int) bool {
	return strings.ToLower(u[i].Title) < strings.ToLower(u[j].Title)
}

func (u userFeedInfos) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
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

	infos := make([]FeedInfo, len(uinfo.Feeds))
	err = datastore.GetMulti(c, uinfo.Feeds, infos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var feeds userFeedInfos
	for i := range infos {
		feeds = append(feeds, userFeedInfo{
			Title:      infos[i].Title,
			LastFetch:  infos[i].LastFetch,
			EncodedKey: uinfo.Feeds[i].Encode(),
		})
	}

	sort.Sort(feeds)

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
		Logout   string
		Title    string
		Link     string
		Errors   []error
		Articles Articles
	}{}

	feedPage.Logout, err = user.LogoutURL(c, "/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.URL.Path == "/" {
		feedPage.Title = "Latest Articles"
		feedPage.Articles, feedPage.Errors = articlesSince(c, uinfo, time.Now().Add(-latestDuration))
	} else if r.URL.Path == "/all" {
		feedPage.Title = "All Articles"
		feedPage.Articles, feedPage.Errors = articlesSince(c, uinfo, time.Time{})
	} else {
		var key *datastore.Key
		var err error
		if key, err = datastore.DecodeKey(r.URL.Path[1:]); err != nil {
			http.NotFound(w, r)
			return
		}

		var f FeedInfo
		if err = datastore.Get(c, key, &f); err != nil {
			err = fmt.Errorf("%s: failed to load from the datastore: %s", key.StringID(), err.Error())
			feedPage.Errors = []error{err}

		} else if err := f.ensureFresh(c); err != nil {
			err = fmt.Errorf("%s: failed to refresh: %s", f.Url, err.Error())
			feedPage.Errors = []error{err}

		} else {
			feedPage.Title = f.Title
			feedPage.Link = f.Link
			feedPage.Articles, err = f.articlesSince(c, time.Time{})
			if err != nil {
				feedPage.Errors = []error{err}
			}
		}
	}

	c.Debugf("%d articles\n", len(feedPage.Articles))
	sort.Sort(feedPage.Articles)

	if err := templates.ExecuteTemplate(w, "feed.html", feedPage); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func articlesSince(c appengine.Context, uinfo UserInfo, t time.Time) (articles Articles, errs []error) {
	for _, key := range uinfo.Feeds {
		var f FeedInfo
		if err := datastore.Get(c, key, &f); err != nil {
			err = fmt.Errorf("%s: failed to load from the datastore: %s", key.StringID(), err.Error())
			errs = append(errs, err)
			continue
		}
		if err := f.ensureFresh(c); err != nil {
			err = fmt.Errorf("%s: failed to refresh: %s", f.Url, err.Error())
			errs = append(errs, err)
			continue
		}
		as, err := f.articlesSince(c, t)
		if err != nil {
			err = fmt.Errorf("%s: failed to read articles: %s", f.Url, err.Error())
			errs = append(errs, err)
			continue
		}
		articles = append(articles, as...)
	}
	return
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	c := appengine.NewContext(r)

	// Check tha the feed is even valid, and put it in the datastore.
	url := r.FormValue("url")
	f, err := checkUrl(c, url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := subscribe(c, f); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/list", http.StatusFound)
}

type Outline struct {
	XmlURL   string     `xml:"xmlUrl,attr"`
	Outlines []*Outline `xml:"outline"`
}

func handleOpml(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	c := appengine.NewContext(r)

	f, _, err := r.FormFile("opml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var b struct {
		Body Outline `xml:"body"`
	}
	err = xml.NewDecoder(f).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	urls := opmlWalk(&b.Body, nil)

	c.Debugf("Got %d URLs from OPML", len(urls))

	for _, url := range urls {
		c.Debugf("opml %s", url)
		f, err := checkUrl(c, url)
		if err != nil {
			http.Error(w, "failed to check URL "+url+": "+err.Error(), http.StatusInternalServerError)
			return
		}

		if err = subscribe(c, f); err != nil {
			http.Error(w, "failed to subscribe "+url+": "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/list", http.StatusFound)
}

func opmlWalk(r *Outline, urls []string) []string {
	if r.XmlURL != "" {
		urls = append(urls, r.XmlURL)
	}
	for _, kid := range r.Outlines {
		urls = append(urls, opmlWalk(kid, nil)...)
	}
	return urls
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
