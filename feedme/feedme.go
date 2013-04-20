package feedme

import (
	"net/http"
	"html/template"
	"strings"

	rss "github.com/zippoxer/RSS-Go"
)

func init () {
	http.HandleFunc("/", root)
}

var feedTemplate = template.Must(template.ParseFiles("tmplt/feed.html"))

func root(w http.ResponseWriter, r *http.Request) {
	feed, err := rss.Get(strings.NewReader(testFeed))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := feedTemplate.Execute(w, feed.Items); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
