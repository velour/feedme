package webfeed

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"log"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"github.com/eaburns/pretty"
)

type Feed struct {
	Title   string
	Link    string
	Updated time.Time
	Entries []Entry
}

type Entry struct {
	Title string
	Link  string
	// Summary is a valid HTML or escaped HTML summary of the entry.
	Summary []byte
	// Contents is the main contents of the entry in valid HTML or escaped HTML.
	Content []byte
	When    time.Time
	ID      string
}

// Read reads a feed from an io.Reader and returns it or an error if one was encountered.
//
// RSS is like the wild west with respect to time. When reading RSS, this
// function may return the non-fatal error ErrBadTime containing the
// first unparsable time encountered.
func Read(r io.Reader) (Feed, error) {
	var f feed
	d := xml.NewDecoder(r)
	d.CharsetReader = charsetReader
	if err := d.Decode(&f); err != nil {
		return Feed{}, err
	}
	if f.Rss.Title != "" {
		return rssFeed(f.Rss)
	}
	return atomFeed(f)
}

func charsetReader(charset string, r io.Reader) (io.Reader, error) {
	if charset == "ISO-8859-1" || charset == "iso-8859-1" {
		return r, nil
	}
	return nil, errors.New("Unsupported character set encoding: " + charset)
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
		Link:    r.link(),
		Updated: updated,
	}

	for _, it := range r.Items {
		when, e := rssTime(it.Updated)
		if err == nil && e != nil {
			err = e
		}
		descr := goodAsText(it.Description)
		if len(descr) == 0 {
			descr = []byte(it.Link)
		}
		content := fixHtml(it.Link, it.Content.Data)
		// If there is no contents, then use the maybe-not-good-as-text description.
		if len(content) == 0 {
			content = it.Description
		}
		ent := Entry{
			Title:   it.Title,
			Link:    it.Link,
			Summary: descr,
			Content: content,
			When:    when,
			ID:      it.GUID,
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
	"2006-01-02T15:04:05-07:00",
	"Mon, 2 Jan 2006 15:04:05 Z",
	"02 Jan 2006 15:04:05 -0700",
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
		Link:    a.link(),
		Updated: a.Updated,
	}

	for _, ent := range a.Entries {
		e := Entry{
			Title:   ent.Title,
			Link:    ent.Link.Href,
			Summary: fixHtml(ent.Link.Href, ent.Summary),
			When:    ent.Updated,
			ID:      ent.ID,
		}
		if len(ent.Content) > 0 {
			e.Content = fixHtml(ent.Link.Href, ent.Content[0].Data())
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
	Links   []atomLink  `xml:"link"`
	Updated time.Time   `xml:"updated"`
	Author  []string    `xml:"author>name"`
	ID      string      `xml:"id"`
	Entries []atomEntry `xml:"entry"`
	Rss     rss         `xml:"channel"`
}

func (f *feed) link() string {
	for _, l := range f.Links {
		if l.Rel == "" || l.Rel == "alternate" {
			return l.Href
		}
	}
	return ""
}

type atomEntry struct {
	Title   string        `xml:"title"`
	Link    atomLink      `xml:"link"`
	ID      string        `xml:"id"`
	Updated time.Time     `xml:"updated"`
	Author  []string      `xml:"author>name"`
	Summary []byte        `xml:"summary"`
	Content []atomContent `xml:"content"`
}

type atomLink struct {
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

type atomContent struct {
	Type     string `xml:"type,attr"`
	Contents []byte `xml:",innerxml"`
}

func (c atomContent) Data() []byte {
	unesc := c.Contents
	if c.Type != "xhtml" {
		unesc = []byte(html.UnescapeString(string(c.Contents)))
	}
	return unesc
}

type rss struct {
	Title       string    `xml:"title"`
	Links       []string  `xml:"link"`
	Description []byte    `xml:"description"`
	Items       []rssItem `xml:"item"`

	// RSS uses its own time format (not understood by the XML parser, because it
	// is apparently a different format from all of the rest of XML in all the land).  We
	// read it as a string and parse it later.

	Updated string `xml:"pubDate"`
}

func (r rss) link() string {
	for _, l := range r.Links {
		if l != "" {
			return l
		}
	}
	return ""
}

type rssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description []byte `xml:"description"`

	// GUID is, unfortunately, an optional field.
	GUID string `xml:"guid"`

	// Content contains <content:encoded>, an extension used by Ars Technica's feeds,
	// and some others.
	Content rssContent `xml:"encoded"`
	Updated string     `xml:"pubDate"`
}

type rssContent struct {
	Data []byte `xml:",chardata"`
}

// FixHtml parses bytes as HTML and returns well-formed HTML if the parse
// was successful, or escaped HTML, if not.
func fixHtml(linkUrl string, wild []byte) (well []byte) {
	n, err := html.Parse(bytes.NewReader(wild))
	if err != nil {
		return []byte(html.EscapeString(string(wild)))
	}

	fixImgs(linkUrl, n)

	defer func() {
		if err := recover(); err == bytes.ErrTooLarge {
			well = []byte(html.EscapeString(string(wild)))
		} else if err != nil {
			panic(err)
		}
	}()
	buf := bytes.NewBuffer(make([]byte, 0, len(wild)*2))
	if err := html.Render(buf, n); err != nil {
		return []byte(html.EscapeString(string(wild)))
	}

	well = buf.Bytes()
	openBody := []byte("<body>")
	i := bytes.Index(well, openBody)
	if i < 0 {
		return []byte(html.EscapeString(string(wild)))
	}
	well = well[i+len(openBody):]

	closeBody := []byte("</body>")
	i = bytes.Index(well, closeBody)
	if i < 0 {
		return []byte(html.EscapeString(string(wild)))
	}
	return well[:i]
}

// FixImgs makes the source URLs of image absolute given to the link URL.
// This fixes feeds that specify their images using relative URLs.
func fixImgs(link string, n *html.Node) {
	u, err := url.Parse(link)
	if err != nil {
		return
	}

	walk(n, func(n *html.Node) {
		if n.Type != html.ElementNode || n.DataAtom != atom.Img {
			return
		}
		for i, a := range n.Attr {
			if a.Key != "src" || len(a.Val) == 0 {
				continue
			}
			src, err := u.Parse(a.Val)
			if err != nil {
				continue
			}
			a.Val = src.String()
			n.Attr[i] = a
		}
	})
}

func walk(root *html.Node, f func(*html.Node)) {
	f(root)
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		walk(c, f)
	}
}

// GoodAsText applies some heuristics to make the data look good when displayed
// as simple text. For example, if the data is escaped HTML then other heuristics are
// applied to remove the HTML. For example if the data contains an HTML image tag,
// goodAsText will return the alt text. If nothing good is found then an empty slice is
// returned.
func goodAsText(d []byte) []byte {
	unesc := html.UnescapeString(string(d))
	nodes, err := html.ParseFragment(strings.NewReader(unesc), bodyNode)
	if err != nil {
		log.Printf("failed to parse [%s] as HTML: %v", unesc, err)
		return d
	}

	var buf bytes.Buffer
	for _, root := range nodes {
		walk(root, func(n *html.Node) {
			if n.Type == html.TextNode {
				buf.WriteString(strings.TrimSpace(n.Data))
				return
			}

			if n := buf.Len(); n > 0 && buf.Bytes()[n-1] != ' ' {
				buf.WriteString(" ")
			}
			if n.DataAtom == atom.Img {
				if alt := altTextOrEmpty(n); alt != "" {
					buf.WriteString(alt)
				}
			}
		})
	}
	return buf.Bytes()
}

// AltTextOrEmpty returns the alt text of an image node wrapped in "[" and "]".
// If there is no alt text, the empty string is returned. The node must be an IMG
// node.
func altTextOrEmpty(n *html.Node) string {
	if n.Type != html.ElementNode || n.DataAtom != atom.Img {
		panic("expected an img node")
	}
	for _, attr := range n.Attr {
		if attr.Key == "alt" {
			return "[" + attr.Val + "]"
		}
	}
	return ""
}

var bodyNode = getBodyNode()

// GetBodyNode returns an BODY node nested within an HTML node.
func getBodyNode() *html.Node {
	ns, err := html.ParseFragment(strings.NewReader("<html><body></body></html>"), nil)
	if err != nil {
		panic("error generating context")
	}
	if len(ns) == 0 {
		panic("no nodes generating context")
	}
	h := ns[0]
	if h.Type != html.ElementNode || h.DataAtom != atom.Html {
		panic("expected an HTML node, got " + pretty.String(h))
	}
	b := h.LastChild
	if b.Type != html.ElementNode || b.DataAtom != atom.Body {
		panic("expected a BODY node, got " + pretty.String(b))
	}
	return b
}
