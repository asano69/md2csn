package convert

import "testing"

func TestConvertImages(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"basic", "![a](https://x.com/a.png)\n", "[https://x.com/a.png]\n"},
		{"empty alt", "![](https://x.com/a.png)\n", "[https://x.com/a.png]\n"},
		{"title is dropped", "![a](https://x.com/a.png \"t\")\n", "[https://x.com/a.png]\n"},
		{"relative path is kept", "![a](./a.png)\n", "[./a.png]\n"},
		{"url without an extension is kept", "![a](https://x.com/a)\n", "[https://x.com/a]\n"},
		{"inside a sentence", "x ![a](https://x.com/a.png) y\n", "x [https://x.com/a.png] y\n"},
		{"in a list item", "- ![a](https://x.com/a.png)\n", "\t[https://x.com/a.png]\n"},
		{"not decorated in bold", "**![a](https://x.com/a.png)**\n", "[https://x.com/a.png]\n"},
		{"kept in a heading", "## ![a](https://x.com/a.png)\n", "[** [https://x.com/a.png]]\n"},
		{"kept in a table cell", "| a |\n| - |\n| ![x](https://x.com/a.png) |\n", "table:_\n\ta\n\t[https://x.com/a.png]\n"},
		{"image in an external link", "[![a](https://x.com/a.png)](https://example.com)\n", "[https://example.com https://x.com/a.png]\n"},
		{"image in a relative link", "[![a](https://x.com/a.png)](./page)\n", "[https://x.com/a.png]\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert([]byte(tt.in), ""); got != tt.want {
				t.Errorf("Convert(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
