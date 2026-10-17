package ast // import "github.com/yuin/goldmark/v2/ast"

type BaseBlock struct {
	BaseNode

	// Has unexported fields.
}
    A BaseBlock struct implements the Node interface partially.

func (b *BaseBlock) AppendSource(seg text.Segment)
func (n *BaseBlock) Attribute(name string) (text.MultiLineValue, bool)
func (n *BaseBlock) Attributes() []Attribute
func (b *BaseBlock) HasBlankPreviousLines() bool
func (n *BaseBlock) NextSibling() Node
func (n *BaseBlock) Parent() Node
func (n *BaseBlock) PreviousSibling() Node
func (n *BaseBlock) RemoveAttributes()
func (n *BaseBlock) SetAttribute(name string, value text.MultiLineValue)
func (b *BaseBlock) SetBlankPreviousLines(v bool)
func (n *BaseBlock) SetNextSibling(v Node)
func (n *BaseBlock) SetParent(v Node)
func (n *BaseBlock) SetPreviousSibling(v Node)
func (b *BaseBlock) SetSource(v []text.Segment)
func (b *BaseBlock) Source() []text.Segment
