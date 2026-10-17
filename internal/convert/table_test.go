package convert

import "testing"

func TestConvertTables(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"basic", "| Name | Age |\n| --- | --- |\n| Alice | 11 |\n", "table:_\n\tName\tAge\n\tAlice\t11\n"},
		{"several rows", "| a | b |\n| - | - |\n| 1 | 2 |\n| 3 | 4 |\n", "table:_\n\ta\tb\n\t1\t2\n\t3\t4\n"},
		{"header only", "| a | b |\n| - | - |\n", "table:_\n\ta\tb\n"},
		{"alignment is dropped", "| a | b |\n|:-:|--:|\n| 1 | 2 |\n", "table:_\n\ta\tb\n\t1\t2\n"},
		{"wikilink is converted", "| a | b |\n| - | - |\n| [[x]] | y |\n", "table:_\n\ta\tb\n\t[x]\ty\n"},
		{"decoration is dropped", "| a | b |\n| - | - |\n| **x** | y |\n", "table:_\n\ta\tb\n\tx\ty\n"},
		{"empty cell", "| a | b |\n| - | - |\n|  | y |\n", "table:_\n\ta\tb\n\t\ty\n"},
		{"between paragraphs", "text\n\n| a |\n| - |\n| 1 |\n\nmore\n", "text\n\ntable:_\n\ta\n\t1\n\nmore\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Convert([]byte(tt.in), ""); got != tt.want {
				t.Errorf("Convert(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}
