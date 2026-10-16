// Package convert converts Markdown text to Cosense notation.
package convert

import (
	"iter"

	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/parser"
)

// Convert converts Markdown source to Cosense notation.
func Convert(source []byte) string {
	doc := parser.New().Parse(source)
	r := &renderer{source: source, stars: headingStars(doc)}

	out := r.renderBlocks(doc)
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
