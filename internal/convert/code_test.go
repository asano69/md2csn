package convert

import "testing"

func TestConvertCodeBlocks(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"fenced with language", "```go\nx := 1\ny := 2\n```\n", "code: go\n\tx := 1\n\ty := 2\n"},
		{"fenced without language", "```\nx\n```\n", "code:\n\tx\n"},
		{"only the first word of the info string", "```go title=a\nx\n```\n", "code: go\n\tx\n"},
		{"blank line inside is kept", "```\na\n\nb\n```\n", "code:\n\ta\n\t\n\tb\n"},
		{"leading spaces are kept", "```\n  a\n```\n", "code:\n\t  a\n"},
		{"indented code block", "    x\n    y\n", "code:\n\tx\n\ty\n"},
		{"between paragraphs", "a\n\n```\nx\n```\n\nb\n", "a\n\ncode:\n\tx\n\nb\n"},
		{"empty block", "```\n```\n", "code:\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert([]byte(tt.in), ""); got != tt.want {
				t.Errorf("Convert(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
