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
	Feeds []*datastore.Key
}

func init() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/feeds", handleFeeds)
	http.HandleFunc("/add", handleAdd)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" || r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	c := appengine.NewContext(r)

	uinfo, err := userInfo(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := rootTemplate.Execute(w, uinfo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleFeeds(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	uinfo, err := userInfo(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var feed Feed
	if num := r.URL.Query().Get("n"); num == "" {
		feed, err = getFeeds(c, uinfo.Feeds)
	} else {
		var n int
		n, err = strconv.Atoi(num)
		if err != nil || n < 0 || n >= len(uinfo.Feeds) {
			http.NotFound(w, r)
			return
		}
		feed, err = getFeedByKey(c, uinfo.Feeds[n])
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

func handleAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	c := appengine.NewContext(r)

	// Check tha the feed is even valid, and put it in the datastore.
	url := r.FormValue("url")
	if _, err := getFeedByUrl(c, url); err != nil {
		c.Errorf("fetchFede failed: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	u, err := userInfo(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := u.addFeed(c, urlKey(c, url)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

// AddFeed adds a feed to the user's feed list if it is not already there.
// The feed must already be in the datastore.
func (u UserInfo) addFeed(c appengine.Context, feedKey *datastore.Key) error {
	return datastore.RunInTransaction(c, func(c appengine.Context) error {
		for _, k := range u.Feeds {
			if feedKey.Equal(k) {
				return nil
			}
		}

		var f Feed
		if err := datastore.Get(c, feedKey, &f); err != nil{
			return err
		}

		f.Refs++
		if _, err := datastore.Put(c, feedKey, &f); err != nil {
			return err
		}

		u.Feeds = append(u.Feeds, feedKey)
		_, err := datastore.Put(c, userInfoKey(c), &u)
		return err
	}, &datastore.TransactionOptions{ XG: true })
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
