package convert

import "testing"

func TestConvertWikiLinks(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"one link", "see [[Big Bang]]\n", "see [Big Bang]\n"},
		{"suffix after the link", "[[star]]s\n", "[star]s\n"},
		{"two links", "[[a]] and [[b]]\n", "[a] and [b]\n"},
		{"in a list item", "- [[a]]\n", "\t[a]\n"},
		{"empty label is plain text", "[[]]\n", "[[]]\n"},
		{"not closed is plain text", "[[a\n", "[[a\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert([]byte(tt.in), ""); got != tt.want {
				t.Errorf("Convert(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
