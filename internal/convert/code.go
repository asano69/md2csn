package convert

import (
	"strings"

	"github.com/yuin/goldmark/v2/ast"
)

// renderCode renders a code block in Cosense notation: a "code: lang" line
// followed by one tab-indented line per line of code. The language is
// omitted when the block has none, as in an indented code block.
// Blank lines are written as a lone tab so that they do not end the block.
func (r *renderer) renderCode(block *ast.CodeBlock) string {
	header := "code:"
	if lang, _ := block.Language(r.source); lang != "" {
		header += " " + lang
	}
	lines := []string{header}
	// The content ends with a newline, which is not a line of its own.
	body := strings.TrimSuffix(block.Value.Str(r.source), "\n")
	if body != "" {
		for _, line := range strings.Split(body, "\n") {
			lines = append(lines, "\t"+strings.TrimSuffix(line, "\r"))
		}
	}
	return strings.Join(lines, "\n")
}
