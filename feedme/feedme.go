package feedme

import (
	"appengine"
	"appengine/urlfetch"
	"html/template"
	"net/http"
	"sort"

	rss "github.com/zippoxer/RSS-Go"
)

var feedTemplate = template.Must(template.ParseFiles("tmplt/feed.html"))

var feeds = []string{
	"http://feeds.arstechnica.com/arstechnica/security",
	"http://feeds.arstechnica.com/arstechnica/science",
}

func init() {
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)

	feed, err := fetchAggregateFeed(client, feeds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := feedTemplate.Execute(w, feed); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Article struct {
	Feed *rss.Feed
	*rss.Item
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
