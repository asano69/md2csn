package ast // import "github.com/yuin/goldmark/v2/ast"

Package ast defines AST nodes that represent markdown elements.

var KindAutoLink = NewNodeKind("AutoLink")
var KindBlockquote = NewNodeKind("Blockquote")
var KindCodeBlock = NewNodeKind("CodeBlock")
var KindCodeSpan = NewNodeKind("CodeSpan")
var KindDocument = NewNodeKind("Document")
var KindEmphasis = NewNodeKind("Emphasis")
var KindHTMLBlock = NewNodeKind("HTMLBlock")
var KindHeading = NewNodeKind("Heading")
var KindImage = NewNodeKind("Image")
var KindLink = NewNodeKind("Link")
var KindLinkReferenceDefinition = NewNodeKind("LinkReferenceDefinition")
var KindList = NewNodeKind("List")
var KindListItem = NewNodeKind("ListItem")
var KindParagraph = NewNodeKind("Paragraph")
var KindRawHTML = NewNodeKind("RawHTML")
var KindStrong = NewNodeKind("Strong")
var KindText = NewNodeKind("Text")
var KindThematicBreak = NewNodeKind("ThematicBreak")
func IsParagraph(node Node) bool
func Walk(n Node, walker Walker) error
func WithLinkReference(kind ReferenceLinkKind, value text.MultiLineValue) interface{ ... }
func WithLinkTitle(title text.MultiLineValue) interface{ ... }
type Attribute struct{ ... }
type AutoLink struct{ ... }
    func NewAutoLink(destination, label text.SingleLineValue, opts ...AutoLinkOption) *AutoLink
type AutoLinkOption interface{ ... }
    func WithAutoLinkText(text text.SingleLineValue) AutoLinkOption
type BaseBlock struct{ ... }
type BaseInline struct{ ... }
type BaseNode struct{ ... }
type BlockNode interface{ ... }
type Blockquote struct{ ... }
    func NewBlockquote() *Blockquote
type CodeBlock struct{ ... }
    func NewCodeBlock(kind CodeBlockKind, value text.Lines, opts ...CodeBlockOption) *CodeBlock
type CodeBlockKind int
    const CodeBlockKindIndented CodeBlockKind = iota + 1 ...
type CodeBlockOption interface{ ... }
    func WithCodeBlockInfo(info text.SingleLineValue) CodeBlockOption
type CodeSpan struct{ ... }
    func NewCodeSpan(value text.Value) *CodeSpan
type Document struct{ ... }
    func NewDocument() *Document
type Emphasis struct{ ... }
    func NewEmphasis() *Emphasis
type HTMLBlock struct{ ... }
    func NewHTMLBlock(kind HTMLBlockKind) *HTMLBlock
type HTMLBlockKind int
    const HTMLBlockKind1 HTMLBlockKind = iota + 1 ...
type Heading struct{ ... }
    func NewHeading(level int, kind HeadingKind) *Heading
type HeadingKind int
    const HeadingKindATX HeadingKind = iota + 1 ...
type Image struct{ ... }
    func NewImage(destination text.SingleLineValue, opts ...ImageOption) *Image
type ImageOption interface{ ... }
type InlineNode interface{ ... }
type Link struct{ ... }
    func NewLink(destination text.SingleLineValue, opts ...LinkOption) *Link
type LinkOption interface{ ... }
type LinkReferenceDefinition struct{ ... }
    func NewLinkReferenceDefinition(label text.MultiLineValue, destination text.SingleLineValue, ...) *LinkReferenceDefinition
type LinkReferenceDefinitionOption interface{ ... }
type List struct{ ... }
    func NewList(marker byte) *List
type ListItem struct{ ... }
    func NewListItem() *ListItem
type Node interface{ ... }
    func N(node Node, children ...any) Node
type NodeDump struct{ ... }
    func NewNodeDump(node Node, properties map[string]any) *NodeDump
type NodeKind int
    var CurrentKindValue NodeKind
    func NewNodeKind(name string) NodeKind
type Paragraph struct{ ... }
    func NewParagraph() *Paragraph
type PrettyPrintOption func(*ppCfg)
    func WithLevel(level int) PrettyPrintOption
    func WithSource(include bool) PrettyPrintOption
type RawHTML struct{ ... }
    func NewRawHTML(value text.MultiLineValue) *RawHTML
type ReferenceLink struct{ ... }
    func NewReferenceLink(kind ReferenceLinkKind, value text.MultiLineValue) *ReferenceLink
type ReferenceLinkKind int
    const ReferenceLinkKindFull ReferenceLinkKind = iota + 1 ...
type Strong struct{ ... }
    func NewStrong() *Strong
type Text struct{ ... }
    func NewText(value text.SingleLineValue) *Text
type ThematicBreak struct{ ... }
    func NewThematicBreak() *ThematicBreak
type WalkStatus int
    const WalkStop WalkStatus = iota + 1 ...
type Walker func(n Node, entering bool) (WalkStatus, error)
