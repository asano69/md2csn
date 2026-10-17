package convert

import (
	"fmt"
	"strings"

	"github.com/asano69/md2csn/internal/wikilink"
	"github.com/yuin/goldmark/v2/ast"
	extast "github.com/yuin/goldmark/v2/extension/ast"
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
		return fmt.Sprintf("[%s %s]", strings.Repeat("*", r.stars[n.Level]), r.renderInlines(n, stylePlain))
	case *ast.Paragraph:
		return r.renderInlines(n, styleNone)
	case *ast.List:
		return r.renderList(n, 0)
	case *ast.CodeBlock:
		return r.renderCode(n)
	case *extast.Table:
		return r.renderTable(n)
	default:
		// Unsupported blocks only keep the content of their child blocks for now.
		return r.renderBlocks(n)
	}
}

// renderInlines renders the inline children of parent. st holds the
// decorations inherited from the enclosing nodes; stylePlain drops them.
// Links and code spans are never decorated.
func (r *renderer) renderInlines(parent ast.Node, st style) string {
	var sb strings.Builder
	for c := range children(parent) {
		switch n := c.(type) {
		case *ast.Text:
			sb.WriteString(st.wrap(n.Value.Value(r.source)))
			if n.SoftLineBreak() || n.HardLineBreak() {
				sb.WriteByte('\n')
			}
		case *ast.CodeSpan:
			sb.WriteString(n.Value.Value(r.source))
		case *wikilink.Node:
			// [[text]] in Markdown is [text] in Cosense.
			sb.WriteString(st.wrap("[" + r.renderInlines(n, stylePlain) + "]"))
		case *ast.Link:
			sb.WriteString(r.renderLink(n))
		case *ast.Strong:
			sb.WriteString(r.renderInlines(n, st.add(styleBold)))
		case *ast.Emphasis:
			sb.WriteString(r.renderInlines(n, st.add(styleItalic)))
		case *extast.Strikethrough:
			sb.WriteString(r.renderInlines(n, st.add(styleStrike)))
		default:
			// Unsupported inlines only keep their text content for now.
			sb.WriteString(r.renderInlines(n, st))
		}
	}
	return sb.String()
}
