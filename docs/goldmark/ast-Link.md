package ast // import "github.com/yuin/goldmark/v2/ast"

type Link struct {
	BaseInline

	// Destination is a destination(URL) of this link.
	Destination text.SingleLineValue

	// Title is a title of this link.
	Title text.MultiLineValue

	// Reference is a reference of this link. This field is used for reference links.
	// If this link is not a reference link, this field is nil.
	Reference *ReferenceLink
}
    A Link struct represents a link of the Markdown text.

func NewLink(destination text.SingleLineValue, opts ...LinkOption) *Link
func (n *Link) Attribute(name string) (text.MultiLineValue, bool)
func (n *Link) Attributes() []Attribute
func (n *Link) Dump(_ []byte) *NodeDump
func (n *Link) Kind() NodeKind
func (n *Link) NextSibling() Node
func (n *Link) Parent() Node
func (n *Link) PreviousSibling() Node
func (n *Link) RemoveAttributes()
func (n *Link) SetAttribute(name string, value text.MultiLineValue)
func (n *Link) SetNextSibling(v Node)
func (n *Link) SetParent(v Node)
func (n *Link) SetPreviousSibling(v Node)
