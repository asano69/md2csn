package convert

import (
	"strings"

	"github.com/yuin/goldmark/v2/ast"
)

// renderList renders a list as one Cosense line per line of content.
// depth is the number of lists that enclose list.
//
// Bullet items get one tab per list level, counting their own list.
// Ordered items are kept as written ("1. text") because Cosense does not
// parse numbered lists, so only the lists around them add tabs.
func (r *renderer) renderList(list *ast.List, depth int) string {
	var lines []string
	for item := range children(list) {
		lines = append(lines, r.renderItem(item, list.IsOrdered(), depth)...)
	}
	return strings.Join(lines, "\n")
}

// renderItem renders the blocks of one list item. Every line after the first
// is one tab deeper than the first line. Nested lists keep their own depth.
func (r *renderer) renderItem(item ast.Node, ordered bool, depth int) []string {
	indent, prefix := depth+1, ""
	if ordered {
		indent, prefix = depth, r.orderedMarker(item)
	}

	var lines []string
	started := false // whether the first line of the item has been written
	for c := range children(item) {
		if nested, ok := c.(*ast.List); ok {
			if s := r.renderList(nested, depth+1); s != "" {
				lines = append(lines, s)
			}
			continue
		}
		for _, line := range strings.Split(r.renderBlock(c), "\n") {
			if line == "" {
				continue // a blank line would end the list in Cosense
			}
			if started {
				line = strings.Repeat("\t", indent+1) + line
			} else {
				line = strings.Repeat("\t", indent) + prefix + line
				started = true
			}
			lines = append(lines, line)
		}
	}
	return lines
}

// orderedMarker returns the marker written in the source before the content
// of the item, such as "3. ". It returns "" when the marker cannot be found.
func (r *renderer) orderedMarker(item ast.Node) string {
	p, ok := item.FirstChild().(*ast.Paragraph)
	if !ok {
		return ""
	}
	end := p.Pos()
	for end > 0 && (r.source[end-1] == ' ' || r.source[end-1] == '\t') {
		end--
	}
	// source[end-1] is the delimiter ('.' or ')'); digits come before it.
	start := end - 1
	for start > 0 && r.source[start-1] >= '0' && r.source[start-1] <= '9' {
		start--
	}
	if start >= end-1 {
		return ""
	}
	return string(r.source[start:end]) + " "
}
