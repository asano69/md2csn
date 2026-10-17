package convert

import (
	"strings"

	"github.com/yuin/goldmark/v2/ast"
)

// takeTitle returns the title line for the document.
//
// If the document starts with a level-1 heading, its source text is the title
// and the heading is excluded from the rest of the output. The text is kept as
// written: decorations and [[wikilinks]] are not converted. Otherwise fallback
// is the title. A heading elsewhere in the document is never moved.
func (r *renderer) takeTitle(doc ast.Node, fallback string) string {
	h, ok := doc.FirstChild().(*ast.Heading)
	if !ok || h.Level != 1 {
		return fallback
	}
	r.skip = h

	if text := r.sourceText(h); text != "" {
		return text
	}
	return fallback
}

// sourceText returns the source lines of h joined into one line, because a
// title must stay on one line, even for a multi-line setext heading.
func (r *renderer) sourceText(h *ast.Heading) string {
	var lines []string
	for _, seg := range h.Source() {
		if s := strings.TrimSpace(string(seg.Bytes(r.source))); s != "" {
			lines = append(lines, s)
		}
	}
	return strings.Join(lines, " ")
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
