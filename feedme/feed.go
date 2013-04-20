package feedme

import (
	"appengine"
	"appengine/urlfetch"

	rss "github.com/zippoxer/RSS-Go"
)

type Article struct {
	// Origin is the originating feed.
	Origin *rss.Feed
	*rss.Item
}

func (a Article) DateTime() string {
	return a.When.Format("2006-01-02 15:04:05")
}

func (a Article) TimeString() string {
	return a.When.Format("Mon Jan 2 15:04:05 MST 2006")
}

type Feed struct {
	Title string
	Articles []Article
}

func (f Feed) Len() int {
	return len(f.Articles)
}

func (f Feed) Less(i, j int) bool {
	return f.Articles[i].When.After(f.Articles[j].When)
}

func (f Feed) Swap(i, j int) {
	f.Articles[i], f.Articles[j] = f.Articles[j], f.Articles[i]
}

// FetchFeed returns a Feed fetched from the given URL.
func fetchFeed(c appengine.Context, url string) (Feed, error) {
	resp, err := urlfetch.Client(c).Get(url)
	if err != nil {
		return Feed{}, err
	}
	defer resp.Body.Close()

	feed, err := rss.Get(resp.Body)
	if err != nil {
		return Feed{}, err
	}
	as := make([]Article, len(feed.Items))
	for i := range as {
		as[i] = Article{Origin: feed, Item: feed.Items[i]}
	}
	return Feed{ Title: feed.Title, Articles: as }, nil
}

// FetchAll returns an aggregate feed fetched from a set of URLs.
func fetchAll(c appengine.Context, feeds []string) (Feed, error) {
	var as []Article
	for _, url := range feeds {
		f, err := fetchFeed(c, url)
		if err != nil {
			return Feed{}, err
		}
		as = append(as, f.Articles...)
	}
	return Feed{ Title: "All Feeds", Articles: as }, nil
}
