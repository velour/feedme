package feedme

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"html/template"
	"net/http"
	"sort"
	"strconv"
)

var rootTemplate = template.Must(template.ParseFiles("tmplt/root.html"))
var feedTemplate = template.Must(template.ParseFiles("tmplt/feed.html"))

type UserInfo struct {
	Feeds []string
}

func init() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/feeds", handleFeeds)
	http.HandleFunc("/add", handleAdd)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		goToLogin(w, r)
		return
	}

	uinfo, err := userInfo(u, c)
	if err != nil  {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := rootTemplate.Execute(w, uinfo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleFeeds(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		goToLogin(w, r)
		return
	}

	uinfo, err := userInfo(u, c)
	if err != nil  {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var feed Feed
	if num := r.URL.Query().Get("n"); num == "" {
		feed, err = fetchAll(c, uinfo.Feeds)
	} else {
		var n int
		n, err = strconv.Atoi(num)
		if err != nil || n < 0 || n >= len(uinfo.Feeds) {
			http.NotFound(w, r)
			return
		}
		feed, err = fetchFeed(c, uinfo.Feeds[n])
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sort.Sort(feed)

	if err := feedTemplate.Execute(w, feed); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// BUG(eaburns): Add reporting for success and failure
func handleAdd(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		goToLogin(w, r)
		return
	}

	// Check tha the feed is even valid.
	url := r.FormValue("url")
	_, err := fetchFeed(c, url)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	err = datastore.RunInTransaction(c, func(c appengine.Context) error {
		uinfo, err := userInfo(u, c)
		if err != nil {
			return err
		}
		for _, f := range uinfo.Feeds {
			if f == url {
				return nil
			}
		}
		uinfo.Feeds = append(uinfo.Feeds, url)
		_, err = datastore.Put(c, userInfoKey(u, c), &uinfo)
		return err
	}, nil)

	if err != nil {
		c.Errorf("Transaction failed: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

// goToLogin tries to redirect the user to the login page.
func goToLogin(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	url, err := user.LoginURL(c, r.URL.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", url)
	w.WriteHeader(http.StatusFound)
}

// UserInfo returns the UserInfo for the currently logged in user.
// This function assumes that a user is loged in, otherwise it will panic.
func userInfo(u *user.User, c appengine.Context) (UserInfo, error) {
	var uinfo UserInfo
	err := datastore.Get(c, userInfoKey(u, c), &uinfo)
	if err != nil && err != datastore.ErrNoSuchEntity {
		return UserInfo{}, err
	}
	sort.Strings(uinfo.Feeds)
	return uinfo, nil
}

// UserInfoKey returns the key for the current user's UserInfo.
// This function assumes that a user is loged in, otherwise it will panic.
func userInfoKey(u *user.User, c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "User", u.String(), 0, nil)
}
