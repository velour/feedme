package webfeed

import (
	"encoding/xml"
	"io"
	"time"
)

type Feed struct {
	Title   string
	Link    string
	Updated time.Time
	Entries []Entry
}

type Entry struct {
	Title   string
	Link    string
	Summary []byte
	Content []byte
	When    time.Time
}

// Read reads a feed from an io.Reader and returns it or an error if one was encountered.
//
// RSS is like the wild west with respect to time. When reading RSS, this
// function may return the non-fatal error ErrBadTime containing the
// first unparsable time encountered.
func Read(r io.Reader) (Feed, error) {
	var f feed
	if err := xml.NewDecoder(r).Decode(&f); err != nil {
		return Feed{}, err
	}
	if f.Rss.Title != "" {
		return rssFeed(f.Rss)
	}
	return atomFeed(f)
}

// ErrBadTime is a string containing a time that was not parsable.
type ErrBadTime string

func (e ErrBadTime) Error() string {
	return "Unable to parse time: " + string(e)
}

func rssFeed(r rss) (Feed, error) {
	updated, err := rssTime(r.Updated)
	f := Feed{
		Title:   r.Title,
		Link:    r.Link,
		Updated: updated,
	}

	for _, it := range r.Items {
		when, e := rssTime(it.Updated)
		if err == nil && e != nil {
			err = e
		}
		ent := Entry{
			Title:   it.Title,
			Link:    it.Link,
			Summary: it.Description,
			Content: it.Content.Data,
			When:    when,
		}
		f.Entries = append(f.Entries, ent)
	}
	return f, err
}

// RssTimeFormats is a slice of various time formats encountered in the wild.
var rssTimeFormats = []string{
	"Mon, 2 Jan 2006 15:04:05 -0700",
	"Mon, 2 Jan 2006 15:04:05 MST",
	"Mon, 2 Jan 06 15:04:05 -0700",
	"02 January 2006",
}

// RssTime tries parsing a string using a variety of different time formats.
// If the string could not be parsed then the zero time is returned with an ErrBadTime error.
func rssTime(s string) (time.Time, error) {
	if s == "" {
		return time.Time{}, nil
	}

	for _, f := range rssTimeFormats {
		if t, err := time.Parse(f, s); err == nil {
			return t, nil
		}
	}

	return time.Time{}, ErrBadTime(s)
}

func atomFeed(a feed) (Feed, error) {
	f := Feed{
		Title:   a.Title,
		Link:    a.Link.Href,
		Updated: a.Updated,
	}

	for _, ent := range a.Entries {
		e := Entry{
			Title:   ent.Title,
			Link:    ent.Link.Href,
			Summary: ent.Summary,
			When:    ent.Updated,
		}
		if len(ent.Content) > 0 {
			e.Content = ent.Content[0].Data
		}
		f.Entries = append(f.Entries, e)
	}
	return f, nil
}

// Feed is an intermediate representation used to unmarshall the XML;
// it can represent both an Atom feed an an RSS feed.  After unmarshalling
// this information is moved into a more "clean" format: the exported Feed.
type feed struct {
	Title   string      `xml:"title"`
	Link    atomLink    `xml:"link"`
	Updated time.Time   `xml:"updated"`
	Author  []string    `xml:"author>name"`
	Id      string      `xml:"id"`
	Entries []atomEntry `xml:"entry"`
	Rss     rss         `xml:"channel"`
}

type atomEntry struct {
	Title   string        `xml:"title"`
	Link    atomLink      `xml:"link"`
	Id      string        `xml:"id"`
	Updated time.Time     `xml:"updated"`
	Author  []string      `xml:"author>name"`
	Summary []byte        `xml:"summary"`
	Content []atomContent `xml:"content"`
}

type atomLink struct {
	Href string `xml:"href,attr"`
}

type atomContent struct {
	Data []byte `xml:",innerxml"`
}

type rss struct {
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description []byte    `xml:"description"`
	Items       []rssItem `xml:"item"`

	// RSS uses its own time format (not understood by the XML parser, because it
	// is apparently a different format from all of the rest of XML in all the land).  We
	// read it as a string and parse it later.

	Updated string `xml:"pubDate"`
}

type rssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description []byte `xml:"description"`

	// Content contains <content:encoded>, an extension used by Ars Technica's feeds.
	Content rssContent `xml:"content encoded"`
	Updated string     `xml:"pubDate"`
}

type rssContent struct {
	Data []byte `xml:",innerxml"`
}
