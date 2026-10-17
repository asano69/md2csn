package convert

import (
	"fmt"
	"strings"

	"github.com/yuin/goldmark/v2/ast"
)

type renderer struct {
	source []byte
	stars  map[int]int // heading level -> asterisk count
	skip   ast.Node    // block already used as the title; it is not rendered
}

// renderBlocks renders the block children of parent, separated by blank lines.
func (r *renderer) renderBlocks(parent ast.Node) string {
	var blocks []string
	for c := range children(parent) {
		if c == r.skip {
			continue
		}
		if s := r.renderBlock(c); s != "" {
			blocks = append(blocks, s)
		}
	}
	return strings.Join(blocks, "\n\n")
}

func (r *renderer) renderBlock(n ast.Node) string {
	switch n := n.(type) {
	case *ast.Heading:
		return fmt.Sprintf("[%s %s]", strings.Repeat("*", r.stars[n.Level]), r.renderInlines(n))
	case *ast.Paragraph:
		return r.renderInlines(n)
	case *ast.List:
		return r.renderList(n, 0)
	default:
		// Unsupported blocks only keep the content of their child blocks for now.
		return r.renderBlocks(n)
	}
}

// renderInlines renders the inline children of parent as plain text.
func (r *renderer) renderInlines(parent ast.Node) string {
	var sb strings.Builder
	for c := range children(parent) {
		switch n := c.(type) {
		case *ast.Text:
			sb.WriteString(n.Value.Value(r.source))
			if n.SoftLineBreak() || n.HardLineBreak() {
				sb.WriteByte('\n')
			}
		case *ast.CodeSpan:
			sb.WriteString(n.Value.Value(r.source))
		default:
			// Unsupported inlines only keep their text content for now.
			sb.WriteString(r.renderInlines(n))
		}
	}
	return sb.String()
}
