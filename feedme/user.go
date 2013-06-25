package feedme

import (
	"appengine"
	"appengine/datastore"
	"appengine/memcache"
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

		err = fixMissingFieldError(datastore.Get(c, key, &f))
		if err != nil && err != datastore.ErrNoSuchEntity {
			return err
		}

		f.Refs++
		if _, err := datastore.Put(c, key, &f); err != nil {
			return err
		}

		u.Feeds = append(u.Feeds, key)
		return putUser(c, &u)
	}, &datastore.TransactionOptions{XG: true})
	if err != nil {
		return err
	}

	if f.Refs == 1 {
		memcache.Delete(c, mcacheFeedsKey)
		c.Debugf("adding a task to refresh %s\n", key)
		t := taskqueue.NewPOSTTask("/refresh", map[string][]string{"feed": {key.Encode()}})
		_, err = taskqueue.Add(c, t, "")
	}

	return memcache.Gob.Set(c, &memcache.Item{Key: key.StringID(), Object: f})
}

// Unsubscribe removes a feed from the user's feed list.
func unsubscribe(c appengine.Context, feedKey *datastore.Key) error {
	var f FeedInfo
	err := datastore.RunInTransaction(c, func(c appengine.Context) error {
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

		err = fixMissingFieldError(datastore.Get(c, feedKey, &f))
		if err != nil {
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
	if err != nil {
		return err
	}

	if f.Refs <= 0 {
		memcache.Delete(c, mcacheFeedsKey)
		memcache.Delete(c, feedKey.StringID())
		return nil
	}
	return memcache.Gob.Set(c, &memcache.Item{Key: feedKey.StringID(), Object: f})
}

func getUser(c appengine.Context) (UserInfo, error) {
	id := user.Current(c).String()

	var u UserInfo
	if _, err := memcache.Gob.Get(c, id, &u); err == nil {
		return u, nil
	}

	k := datastore.NewKey(c, userKind, id, 0, nil)
	err := datastore.Get(c, k, &u)
	if err != nil && err != datastore.ErrNoSuchEntity {
		return u, err
	}

	err = memcache.Gob.Set(c, &memcache.Item{Key: id, Object: u})
	return u, err
}

func putUser(c appengine.Context, u *UserInfo) error {
	id := user.Current(c).String()
	k := datastore.NewKey(c, userKind, id, 0, nil)
	if _, err := datastore.Put(c, k, u); err != nil {
		return err
	}

	return memcache.Gob.Set(c, &memcache.Item{Key: id, Object: u})
}
