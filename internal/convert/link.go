package convert

import (
	"strings"

	"github.com/yuin/goldmark/v2/ast"
)

// renderLink renders a Markdown link. An external link [text](url) becomes
// [text url] in Cosense. Other links, such as relative paths, keep only
// their text. Decorations in the label are dropped.
// An external link whose label is only an image becomes [url image-url].
func (r *renderer) renderLink(link *ast.Link) string {
	dest := link.Destination.Value(r.source)
	if img, ok := soleImage(link); ok && isExternal(dest) {
		return "[" + dest + " " + img.Destination.Value(r.source) + "]"
	}
	label := r.renderInlines(link, stylePlain)
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
