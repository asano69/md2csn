package wikilink

import (
	"testing"

	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/parser"
)

// labels returns the label of every wikilink in the first paragraph of src.
func labels(src string) []string {
	b := []byte(src)
	doc := parser.New(parser.WithExtensions(Extension)).Parse(b)
	var got []string
	for c := doc.FirstChild().FirstChild(); c != nil; c = c.NextSibling() {
		if w, ok := c.(*Node); ok {
			got = append(got, w.FirstChild().(*ast.Text).Value.Value(b))
		}
	}
	return got
}

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want []string
	}{
		{"one link", "a [[b c]]s d\n", []string{"b c"}},
		{"two links", "[[a]] and [[b]]\n", []string{"a", "b"}},
		{"inside emphasis", "*[[a]]*\n", nil},
		{"normal link is not a wikilink", "[x](u)\n", nil},
		{"empty label", "[[]]\n", nil},
		{"not closed", "[[a\n", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := labels(tt.in)
			if len(got) != len(tt.want) {
				t.Fatalf("labels(%q) = %q, want %q", tt.in, got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("labels(%q) = %q, want %q", tt.in, got, tt.want)
				}
			}
		})
	}
}
