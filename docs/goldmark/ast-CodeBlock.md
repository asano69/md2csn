package ast // import "github.com/yuin/goldmark/v2/ast"

type CodeBlock struct {

	// CodeBlockKind indicates whether this is an indented or fenced code block.
	CodeBlockKind CodeBlockKind

	// Info is the info string of a fenced code block (e.g. language identifier).
	// It is empty for indented code blocks.
	Info text.SingleLineValue

	// Value holds the raw content of this code block for rendering.
	Value text.Lines
	// Has unexported fields.
}
    A CodeBlock struct represents a code block of Markdown text.

func NewCodeBlock(kind CodeBlockKind, value text.Lines, opts ...CodeBlockOption) *CodeBlock
func (n *CodeBlock) AppendChild(_ Node)
func (b *CodeBlock) AppendSource(_ text.Segment)
func (n *CodeBlock) Attribute(name string) (text.MultiLineValue, bool)
func (n *CodeBlock) Attributes() []Attribute
func (n *CodeBlock) ChildCount() int
func (n *CodeBlock) Children() iter.Seq[Node]
func (n *CodeBlock) Dump(_ []byte) *NodeDump
func (n *CodeBlock) FirstChild() Node
func (b *CodeBlock) HasBlankPreviousLines() bool
func (n *CodeBlock) HasChildren() bool
func (n *CodeBlock) InsertAfter(_ Node, _ Node)
func (n *CodeBlock) InsertBefore(_ Node, _ Node)
func (n *CodeBlock) Kind() NodeKind
func (n *CodeBlock) Language(source []byte) (string, bool)
func (n *CodeBlock) LastChild() Node
func (n *CodeBlock) NextSibling() Node
func (n *CodeBlock) OwnerDocument() *Document
func (n *CodeBlock) Parent() Node
func (b *CodeBlock) Pos() int
func (n *CodeBlock) PreviousSibling() Node
func (n *CodeBlock) RemoveAttributes()
func (n *CodeBlock) RemoveChild(_ Node)
func (n *CodeBlock) RemoveChildren()
func (n *CodeBlock) ReplaceChild(_ Node, _ Node)
func (n *CodeBlock) SetAttribute(name string, value text.MultiLineValue)
func (b *CodeBlock) SetBlankPreviousLines(v bool)
func (n *CodeBlock) SetNextSibling(v Node)
func (n *CodeBlock) SetParent(v Node)
func (b *CodeBlock) SetPos(v int)
func (n *CodeBlock) SetPreviousSibling(v Node)
func (b *CodeBlock) SetSource(_ []text.Segment)
func (b *CodeBlock) Source() []text.Segment
