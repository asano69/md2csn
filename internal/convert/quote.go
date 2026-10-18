package convert

import (
	"strings"

	"github.com/yuin/goldmark/v2/ast"
)

// renderQuote renders a blockquote in Cosense notation: one "> text" line per
// line of content. Line breaks inside a paragraph are joined with a space, so
// a paragraph is always one line. No blank line is written between the
// paragraphs, because a blank line would end the quote in Cosense.
// Decorations inside the quote are kept.
func (r *renderer) renderQuote(quote *ast.Blockquote) string {
	var lines []string
	for c := range children(quote) {
		text := r.renderBlock(c)
		if _, ok := c.(*ast.Paragraph); ok {
			text = strings.TrimSpace(strings.ReplaceAll(text, "\n", " "))
		}
		for _, line := range strings.Split(text, "\n") {
			if line != "" {
				lines = append(lines, "> "+line)
			}
		}
	}
	return strings.Join(lines, "\n")
}
