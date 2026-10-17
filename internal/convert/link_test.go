package convert

import "testing"

func TestConvertLinks(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"https link", "[example](https://example.com)\n", "[example https://example.com]\n"},
		{"http link", "[example](http://example.com)\n", "[example http://example.com]\n"},
		{"link inside a sentence", "see [a](https://x.com) now\n", "see [a https://x.com] now\n"},
		{"label with spaces", "[Big Bang](https://example.com/bb)\n", "[Big Bang https://example.com/bb]\n"},
		{"empty label is the URL only", "[](https://example.com)\n", "[https://example.com]\n"},
		{"in a list item", "- [a](https://x.com)\n", "\t[a https://x.com]\n"},
		{"relative link keeps only the text", "[space](<./Space_(physics)>)\n", "space\n"},
		{"wikilink is not an external link", "[[a]]\n", "[a]\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert([]byte(tt.in), ""); got != tt.want {
				t.Errorf("Convert(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
