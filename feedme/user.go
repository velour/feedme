package feedme

import (
	"appengine"
	"appengine/datastore"
	"appengine/taskqueue"
	"appengine/user"
	"fmt"
)

const (
	// MaxFeeds is the maximum allowed number of feeds for a single user.
	maxFeeds = 50

	userKind = "User"
)

type UserInfo struct {
	Feeds []*datastore.Key `datastore:",noindex"`
}

// Subscribe adds a feed to the user's feed list if it is not already there.
func subscribe(c appengine.Context, f FeedInfo) error {
	key := datastore.NewKey(c, feedKind, f.Url, 0, nil)
	err := datastore.RunInTransaction(c, func(c appengine.Context) error {
		u, err := getUser(c)
		if err != nil {
			return err
		}

		if len(u.Feeds) >= maxFeeds {
			return fmt.Errorf("Too many feeds, max is %d", maxFeeds)
		}

		for _, k := range u.Feeds {
			if key.Equal(k) {
				return nil
			}
		}

		if err := datastore.Get(c, key, &f); err != nil && err != datastore.ErrNoSuchEntity {
			return err
		}

		f.Refs++
		if _, err := datastore.Put(c, key, &f); err != nil {
			return err
		}

		u.Feeds = append(u.Feeds, key)
		return putUser(c, &u)
	}, &datastore.TransactionOptions{XG: true})

	if err == nil && f.Refs == 1 {
		c.Debugf("adding a task to refresh %s\n", key)
		t := taskqueue.NewPOSTTask("/refresh", map[string][]string{"feed": {key.Encode()}})
		_, err = taskqueue.Add(c, t, "")
	}
	return err
}

// Unsubscribe removes a feed from the user's feed list.
func unsubscribe(c appengine.Context, feedKey *datastore.Key) error {
	return datastore.RunInTransaction(c, func(c appengine.Context) error {
		u, err := getUser(c)
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
			if err := f.rmArticles(c); err != nil {
				return err
			}
			if err := datastore.Delete(c, feedKey); err != nil {
				return err
			}
		} else if _, err := datastore.Put(c, feedKey, &f); err != nil {
			return err
		}

		u.Feeds = append(u.Feeds[:i], u.Feeds[i+1:]...)
		return putUser(c, &u)
	}, &datastore.TransactionOptions{XG: true})
}

func getUser(c appengine.Context) (UserInfo, error) {
	var uinfo UserInfo
	k := datastore.NewKey(c, userKind, user.Current(c).String(), 0, nil)
	err := datastore.Get(c, k, &uinfo)
	if err == datastore.ErrNoSuchEntity {
		err = nil
	}
	return uinfo, err
}

func putUser(c appengine.Context, u *UserInfo) error {
	k := datastore.NewKey(c, userKind, user.Current(c).String(), 0, nil)
	_, err := datastore.Put(c, k, u)
	return err
}
