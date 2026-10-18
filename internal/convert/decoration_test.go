package convert

import "testing"

func TestConvertDecorations(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"bold", "**a**\n", "[* a]\n"},
		{"italic", "*a*\n", "[/ a]\n"},
		{"italic with underscores", "_a_\n", "[/ a]\n"},
		{"strikethrough", "~~a~~\n", "[- a]\n"},
		{"bold italic", "***a***\n", "[*/ a]\n"},
		{"bold italic strikethrough", "***~~a~~***\n", "[*/- a]\n"},
		{"inside a sentence", "x **a** y\n", "x [* a] y\n"},
		{"whitespace stays outside the bracket", "**a *b* c**\n", "[* a] [*/ b] [* c]\n"},
		{"strikethrough inside bold", "**a ~~b~~**\n", "[* a] [*- b]\n"},
		{"soft line break", "**a\nb**\n", "[* a]\n[* b]\n"},
		{"wikilink is nested", "**[[a]]**\n", "[* [a]]\n"},
		{"link is not decorated", "**[a](https://x.com)**\n", "[a https://x.com]\n"},
		{"code span is not decorated", "**`a`**\n", "`a`\n"},
		{"in a list item", "- **a**\n", "\t[* a]\n"},
		{"dropped in a heading", "## **a**\n", "[** a]\n"},
		{"dropped in a table cell", "| a |\n| - |\n| ~~x~~ |\n", "table:_\n\ta\n\tx\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert([]byte(tt.in), ""); got != tt.want {
				t.Errorf("Convert(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
