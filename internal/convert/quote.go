package convert

import (
	"strings"

	"github.com/yuin/goldmark/v2/ast"
)

// renderQuote renders a blockquote in Cosense notation: one "> text" line per
// line of content. No blank line is written between the paragraphs, because a
// blank line would end the quote in Cosense.
// Decorations inside the quote are kept.
func (r *renderer) renderQuote(quote *ast.Blockquote) string {
	var lines []string
	for c := range children(quote) {
		text := r.renderBlock(c)
		var block []string
		if p, ok := c.(*ast.Paragraph); ok {
			block = r.quoteParagraph(p, text)
		} else {
			block = strings.Split(text, "\n")
		}
		for _, line := range block {
			if line != "" {
				lines = append(lines, "> "+line)
			}
		}
	}
	return strings.Join(lines, "\n")
}

// quoteParagraph splits the rendered text of a paragraph into lines. Lines
// that start with ">" in the source stay separate. A lazy continuation line,
// which has no ">", is joined to the previous line with a space.
// When the rendered lines do not match the source lines, all lines are joined.
func (r *renderer) quoteParagraph(p *ast.Paragraph, text string) []string {
	parts := strings.Split(text, "\n")
	segs := p.Source()
	if len(parts) != len(segs) {
		return []string{strings.Join(parts, " ")}
	}
	var lines []string
	for i, part := range parts {
		part = strings.TrimSpace(part)
		if i > 0 && !r.hasQuoteMarker(segs[i].Start) {
			lines[len(lines)-1] += " " + part
			continue
		}
		lines = append(lines, part)
	}
	return lines
}

// hasQuoteMarker reports whether the line whose content starts at pos begins
// with a blockquote marker.
func (r *renderer) hasQuoteMarker(pos int) bool {
	for pos > 0 && (r.source[pos-1] == ' ' || r.source[pos-1] == '\t') {
		pos--
	}
	return pos > 0 && r.source[pos-1] == '>'
}
