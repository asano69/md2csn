package convert

import "testing"

func TestConvertLists(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"bullets", "- a\n- b\n", "\ta\n\tb\n"},
		{"star bullets", "* a\n* b\n", "\ta\n\tb\n"},
		{"plus bullets", "+ a\n+ b\n", "\ta\n\tb\n"},
		{"changing the marker starts a new list", "* a\n+ b\n", "\ta\n\n\tb\n"},
		{"nested bullets", "- a\n  - b\n    - c\n", "\ta\n\t\tb\n\t\t\tc\n"},
		{"soft line break is joined with a space", "- aaa\n  bbb\n", "\taaa bbb\n"},
		{"soft line break in an ordered item", "1. a\n   b\n", "1. a b\n"},
		{"second paragraph goes one tab deeper", "- a\n\n  b\n", "\ta\n\t\tb\n"},
		{"loose list has no blank lines", "- a\n\n- b\n", "\ta\n\tb\n"},
		{"blank line before and after", "text\n\n- a\n\nmore\n", "text\n\n\ta\n\nmore\n"},
		{"ordered is kept", "1. apple\n1. banana\n", "1. apple\n1. banana\n"},
		{"ordered numbers are kept", "3. a\n4. b\n", "3. a\n4. b\n"},
		{"ordered inside bullet", "- 1. tomato\n", "\t1. tomato\n"},
		{"ordered nested in ordered", "1. a\n   1. b\n", "1. a\n\t1. b\n"},
		{"bullet nested in ordered", "1. a\n   - b\n", "1. a\n\t\tb\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert([]byte(tt.in), ""); got != tt.want {
				t.Errorf("Convert(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
