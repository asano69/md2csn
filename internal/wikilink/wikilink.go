// Package wikilink adds the [[text]] syntax to the goldmark parser.
package wikilink

import (
	"bytes"

	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/parser"
	"github.com/yuin/goldmark/v2/text"
	"github.com/yuin/goldmark/v2/util"
)

// Kind is the NodeKind of Node.
var Kind = ast.NewNodeKind("WikiLink")

// Node is an inline [[text]] link. Its only child is the Text between the
// brackets.
type Node struct {
	ast.BaseInline
}

// Dump implements ast.Node.
func (n *Node) Dump(_ []byte) *ast.NodeDump {
	return ast.NewNodeDump(n, nil)
}

// Kind implements ast.Node.
func (n *Node) Kind() ast.NodeKind {
	return Kind
}

type wikiLinkParser struct{}

func (p *wikiLinkParser) Trigger() []byte {
	return []byte{'['}
}

func (p *wikiLinkParser) Parse(_ ast.Node, block text.Reader, _ parser.Context) ast.Node {
	line, seg := block.PeekLine()
	if !bytes.HasPrefix(line, []byte("[[")) {
		return nil
	}
	n := bytes.Index(line[2:], []byte("]]"))
	if n <= 0 {
		return nil // not closed on this line, or empty
	}
	block.Advance(n + 4)

	link := &Node{}
	link.Init(link)
	label := text.NewSegment(seg.Start+2, seg.Start+2+n)
	link.AppendChild(ast.NewText(text.NewSingleLineValueFromSegment(label, block.Decoder())))
	return link
}

type extension struct{}

// Extension parses [[text]] into Node.
var Extension parser.Extension = extension{}

func (extension) ParserOptions(_ *parser.Config) []parser.Option {
	// The priority is higher than the link parser (200), which also starts with '['.
	return []parser.Option{
		parser.WithInlineParsers(util.Prioritized[parser.InlineParser](&wikiLinkParser{}, 150)),
	}
}
