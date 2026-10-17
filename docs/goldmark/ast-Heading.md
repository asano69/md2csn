package ast // import "github.com/yuin/goldmark/v2/ast"

type Heading struct {
	BaseBlock
	// Level is the level of this heading.
	// This value is between 1 and 6.
	Level int
	// HeadingKind indicates whether this is an ATX or Setext heading.
	HeadingKind HeadingKind
}
    A Heading struct represents headings like SetextHeading and ATXHeading.

func NewHeading(level int, kind HeadingKind) *Heading
func (n *Heading) Attribute(name string) (text.MultiLineValue, bool)
func (n *Heading) Attributes() []Attribute
func (n *Heading) Dump(_ []byte) *NodeDump
func (n *Heading) Kind() NodeKind
func (n *Heading) NextSibling() Node
func (n *Heading) Parent() Node
func (n *Heading) PreviousSibling() Node
func (n *Heading) RemoveAttributes()
func (n *Heading) SetAttribute(name string, value text.MultiLineValue)
func (n *Heading) SetNextSibling(v Node)
func (n *Heading) SetParent(v Node)
func (n *Heading) SetPreviousSibling(v Node)
