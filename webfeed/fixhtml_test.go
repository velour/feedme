package webfeed

import "testing"

func TestFixHtml(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{"", ""},
		{"<div></div>", "<div></div>"},
		{"<div></div></div>", "<div></div>"},
		{"</foo>", ""},
		{"</foo>bar<div>baz</div>", "bar<div>baz</div>"},
		{`<img src="foo.png"/>`, `<img src="http://test.com/foo.png"/>`},
		{`<img src="http://otherdomain.com/foo.png"/>`, `<img src="http://otherdomain.com/foo.png"/>`},
	}

	for _, test := range tests {
		o := string(fixHtml("http://test.com/", []byte(test.in)))
		if o != test.out {
			t.Errorf(`fixHtml("http://test.com/", "%s")="%s", want "%s"`, test.in, o, test.out)
		}
	}
}

func TestGoodAsText(t *testing.T) {
	tests := []struct {
		in, out string
	}{
		{"normal text", "normal text"},
		{"stylized<b>text</b>", "stylized text"},
		{"stylized       <b>text</b>", "stylized text"},
		{"stylized &lt;b&gt;text&lt;/b&gt;", "stylized text"},
		{`<img alt="alternate"/><img alt="text"/>`, "[alternate] [text]"},

		// There's no alt text here, so we should strip the entire thing.
		{`&lt;img src=&quot;http://backcomic.com/s/Back_015_rq9h6h.jpg&quot;&gt;&lt;img src=&quot;http://backcomic.com/s/Back_016_h7b8q5.jpg&quot;&gt;`, ""},
	}
	for _, test := range tests {
		in := []byte(test.in)
		if got := string(goodAsText(in)); test.out != got {
			t.Errorf(`goodAsText(%s)="%s", want "%s"`, test.in, got, test.out)
		}
	}
}
