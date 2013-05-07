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
	}

	for _, test := range tests {
		o := string(fixHtml([]byte(test.in)))
		if o != test.out {
			t.Errorf("Expected [%s] to fix to [%s], but got [%s]", test.in, test.out, o)
		}
	}
}
