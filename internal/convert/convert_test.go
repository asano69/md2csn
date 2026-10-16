package convert

import "testing"

func TestConvertHeadings(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"h1 only", "# a\n", "[** a]\n"},
		{"h1 and h2", "# a\n\n## b\n", "[*** a]\n\n[** b]\n"},
		{"h2 and h3", "## a\n\n### b\n", "[*** a]\n\n[** b]\n"},
		{"h1 to h3", "# a\n\n## b\n\n### c\n", "[**** a]\n\n[*** b]\n\n[** c]\n"},
		{"skipped level", "# a\n\n### b\n", "[*** a]\n\n[** b]\n"},
		{"heading and paragraph", "# a\n\ntext\nmore\n", "[** a]\n\ntext\nmore\n"},
		{"setext heading", "a\n===\n", "[** a]\n"},
		{"empty", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert([]byte(tt.in)); got != tt.want {
				t.Errorf("Convert(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
