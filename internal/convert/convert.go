// Package convert converts Markdown text to Cosense notation.
package convert

import (
	"iter"

	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/parser"
)

// Convert converts Markdown source to Cosense notation.
//
// When fileTitle is not empty, the output starts with a title line: the text
// of a leading level-1 heading, or fileTitle if there is none. The title is
// followed directly by the body, without a blank line.
func Convert(source []byte, fileTitle string) string {
	doc := parser.New().Parse(source)
	r := &renderer{source: source}

	var title string
	if fileTitle != "" {
		title = r.takeTitle(doc, fileTitle)
	}
	r.stars = headingStars(doc, r.skip)

	out := withTitle(title, r.renderBlocks(doc))
	if out == "" {
		return ""
	}
	return out + "\n"
}

// children iterates over the direct children of n.
func children(n ast.Node) iter.Seq[ast.Node] {
	return func(yield func(ast.Node) bool) {
		for c := n.FirstChild(); c != nil; c = c.NextSibling() {
			if !yield(c) {
				return
			}
		}
	}
}
