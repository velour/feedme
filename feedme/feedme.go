package feedme

import (
	"appengine"
	"appengine/datastore"
	"appengine/urlfetch"
	"appengine/user"
	"html/template"
	"net/http"
	"sort"

	rss "github.com/zippoxer/RSS-Go"
)

var feedTemplate = template.Must(template.ParseFiles("tmplt/feed.html"))

type UserInfo struct {
	Feeds []string
}

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/add", addFeed)
}

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		url, err := user.LoginURL(c, r.URL.String())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusFound)
		return
	}

	uinfo, err := userInfo(c)
	if err != nil  {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := urlfetch.Client(c)
	feed, err := fetchAggregateFeed(client, uinfo.Feeds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := feedTemplate.Execute(w, feed); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// BUG(eaburns): Add reporting for success and failure
func addFeed(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	u := user.Current(c)
	if u == nil {
		// Not logged in, go to the root page for login redirection.
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	client := urlfetch.Client(c)
	url := r.FormValue("url")
	_, err := fetchFeed(client, url)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	err = datastore.RunInTransaction(c, func(c appengine.Context) error {
		uinfo, err := userInfo(c)
		if err != nil {
			return err
		}
		for _, f := range uinfo.Feeds {
			if f == url {
				return nil
			}
		}
		uinfo.Feeds = append(uinfo.Feeds, url)
		_, err = datastore.Put(c, userInfoKey(c), &uinfo)
		return err
	}, nil)

	if err != nil {
		c.Errorf("Transaction failed: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
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

type Article struct {
	Feed *rss.Feed
	*rss.Item
}

func (a Article) DateTime() string {
	return a.When.Format("2006-01-02 15:04:05")
}

func (a Article) TimeString() string {
	return a.When.Format("Mon Jan 2 15:04:05 MST 2006")
}

type Feed []Article

func (a Feed) Len() int {
	return len(a)
}

func (a Feed) Less(i, j int) bool {
	return a[i].When.After(a[j].When)
}

func (a Feed) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func fetchFeed(client *http.Client, url string) (Feed, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	feed, err := rss.Get(resp.Body)
	if err != nil {
		return nil, err
	}
	f := make(Feed, len(feed.Items))
	for i := range f {
		f[i] = Article{Feed: feed, Item: feed.Items[i]}
	}
	return f, nil
}

func fetchAggregateFeed(client *http.Client, feeds []string) (Feed, error) {
	var feed Feed
	for _, url := range feeds {
		f, err := fetchFeed(client, url)
		if err != nil {
			return nil, err
		}
		feed = append(feed, f...)
	}
	sort.Sort(feed)
	return feed, nil
}
