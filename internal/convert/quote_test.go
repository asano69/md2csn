package convert

import "testing"

func TestConvertQuotes(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"lazy continuation is joined", "> aaa\nbbb\n", "> aaa bbb\n"},
		{"quoted lines stay separate", "> a\n> b\n", "> a\n> b\n"},
		{"lazy line after quoted lines", "> a\n> b\nc\n", "> a\n> b c\n"},
		{"decoration is kept", "> **a** b\n", "> [* a] b\n"},
		{"decoration across quoted lines", "> **a\n> b**\n", "> [* a]\n> [* b]\n"},
		{"wikilink is converted", "> [[a]]\n", "> [a]\n"},
		{"paragraphs have no blank line", "> a\n>\n> b\n", "> a\n> b\n"},
		{"nested quote", "> a\n>\n> > b\n", "> a\n> > b\n"},
		{"nested quoted lines stay separate", "> > a\n> > b\n", "> > a\n> > b\n"},
		{"between paragraphs", "x\n\n> a\n\ny\n", "x\n\n> a\n\ny\n"},
		{"list in a quote", "> - a\n", "> \ta\n"},
		{"empty quote", ">\n", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert([]byte(tt.in), ""); got != tt.want {
				t.Errorf("Convert(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
