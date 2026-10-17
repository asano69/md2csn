package convert

import (
	"strings"

	"github.com/yuin/goldmark/v2/ast"
)

// renderLink renders a Markdown link. An external link [text](url) becomes
// [text url] in Cosense. Other links, such as relative paths, keep only
// their text.
func (r *renderer) renderLink(link *ast.Link) string {
	label := r.renderInlines(link)
	dest := link.Destination.Value(r.source)
	if !isExternal(dest) {
		return label
	}
	if label == "" {
		return "[" + dest + "]"
	}
	return "[" + label + " " + dest + "]"
}

// isExternal reports whether dest is an http or https URL.
func isExternal(dest string) bool {
	return strings.HasPrefix(dest, "http://") || strings.HasPrefix(dest, "https://")
}
