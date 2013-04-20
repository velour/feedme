package feedme

import (
	"net/http"
	"html/template"
	"strings"	
	"io"
	"sort"

	rss "github.com/zippoxer/RSS-Go"
)

var feedTemplate = template.Must(template.ParseFiles("tmplt/feed.html"))

func init () {
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	articles, err := readFeeds(feeds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := feedTemplate.Execute(w, articles); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func readFeeds(feeds []string) (Articles, error) {
	var articles Articles

	for _, feed := range feeds {
		as, err := readArticles(strings.NewReader(feed))
		if err != nil {
			return nil, err
		}
		articles = append(articles, as...)
	}

	sort.Sort(articles)
	return articles, nil
}

type Article struct {
	Feed *rss.Feed
	*rss.Item
}

type Articles []Article

func (a Articles) Len() int {
	return len(a)
}

func (a Articles) Less(i, j int) bool {
	return a[i].When.After(a[j].When)
}

func (a Articles) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func readArticles(in io.Reader) (Articles, error) {
	feed, err := rss.Get(in)
	if err != nil {
		return nil, err
	}
	articles := make([]Article, len(feed.Items))
	for i := range articles {
		articles[i] = Article{ Feed: feed, Item: feed.Items[i] }
	}
	return articles, nil
}
