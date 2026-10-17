package convert

import "testing"

func TestConvertTitle(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"h1 becomes the title", "# a\n\ntext\n", "a\ntext\n"},
		{"setext h1 becomes the title", "a\n===\n\ntext\n", "a\ntext\n"},
		{"only h1", "# a\n", "a\n"},
		{"decoration is kept", "# **Bold** title\n\ntext\n", "**Bold** title\ntext\n"},
		{"wikilink is kept", "# [[a]] b\n\ntext\n", "[[a]] b\ntext\n"},
		{"multi-line setext title is one line", "a [[b]]\nc\n===\n", "a [[b]] c\n"},
		{"wikilink in h2 is nested", "# a\n\n## [[b]]\n", "a\n[** [b]]\n"},
		{"levels are recounted without the title", "# a\n\n## b\n\n### c\n", "a\n[*** b]\n\n[** c]\n"},
		{"no heading inserts the file name", "text\n", "f\ntext\n"},
		{"non-h1 first inserts the file name", "## x\n\ntext\n", "f\n[** x]\n\ntext\n"},
		{"h1 that is not first is not moved", "text\n\n# a\n", "f\ntext\n\n[** a]\n"},
		{"empty h1 inserts the file name", "#\n\ntext\n", "f\ntext\n"},
		{"empty input", "", "f\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert([]byte(tt.in), "f"); got != tt.want {
				t.Errorf("Convert(%q, \"f\") = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
