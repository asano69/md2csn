package convert

import "testing"

func TestConvertCodeSpans(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"basic", "`a`\n", "`a`\n"},
		{"inside a sentence", "x `a` y\n", "x `a` y\n"},
		{"in a list item", "- `a`\n", "\t`a`\n"},
		{"not decorated in bold", "**`a`**\n", "`a`\n"},
		{"not decorated in strikethrough", "~~`a`~~\n", "`a`\n"},
		{"kept in a heading", "## `a`\n", "[** `a`]\n"},
		{"kept in a table cell", "| a |\n| - |\n| `x` |\n", "table:_\n\ta\n\t`x`\n"},
		{"kept in a link label", "[`a`](https://x.com)\n", "[`a` https://x.com]\n"},
		{"backtick inside is kept", "``a`b``\n", "`a`b`\n"},
		{"line break becomes a space", "`a\nb`\n", "`a b`\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert([]byte(tt.in), ""); got != tt.want {
				t.Errorf("Convert(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
