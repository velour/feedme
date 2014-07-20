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
			t.Errorf(`fixHtml("http://test.com/", %s)=%s, want %s`, test.in, o, test.out)
		}
	}
}
