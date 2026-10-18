package convert

import "github.com/yuin/goldmark/v2/ast"

// renderImage renders an image in Cosense notation: [url].
// The alt text and the title are dropped because Cosense has no place for
// them. The destination is written as is, even for a relative path.
// Decorations are never applied to it.
func (r *renderer) renderImage(img *ast.Image) string {
	return "[" + img.Destination.Value(r.source) + "]"
}

// soleImage returns the image when it is the only child of link, as in
// [![alt](image-url)](page-url).
func soleImage(link *ast.Link) (*ast.Image, bool) {
	if link.ChildCount() != 1 {
		return nil, false
	}
	img, ok := link.FirstChild().(*ast.Image)
	return img, ok
}
