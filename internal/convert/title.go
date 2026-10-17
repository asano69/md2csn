package convert

import (
	"strings"

	"github.com/yuin/goldmark/v2/ast"
)

// takeTitle returns the title line for the document.
//
// If the document starts with a level-1 heading, its plain text is the title
// and the heading is excluded from the rest of the output. Otherwise fallback
// is the title. A heading elsewhere in the document is never moved.
func (r *renderer) takeTitle(doc ast.Node, fallback string) string {
	h, ok := doc.FirstChild().(*ast.Heading)
	if !ok || h.Level != 1 {
		return fallback
	}
	r.skip = h

	// A title must stay on one line, even for a multi-line setext heading.
	text := strings.ReplaceAll(r.renderInlines(h), "\n", " ")
	if text == "" {
		return fallback
	}
	return text
}

// withTitle puts title on the first line, directly followed by body.
func withTitle(title, body string) string {
	if title == "" {
		return body
	}
	if body == "" {
		return title
	}
	return title + "\n" + body
}
