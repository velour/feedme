package feedme

import (
	"net/http"
	"html/template"
	"strings"	
	"io"

	rss "github.com/zippoxer/RSS-Go"
)

func init () {
	http.HandleFunc("/", root)
}

var feedTemplate = template.Must(template.ParseFiles("tmplt/feed.html"))

func root(w http.ResponseWriter, r *http.Request) {
	articles, err := readArticles(strings.NewReader(testFeed))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := feedTemplate.Execute(w, articles); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type Article struct {
	Feed *rss.Feed
	*rss.Item
}

func readArticles(in io.Reader) ([]Article, error) {
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