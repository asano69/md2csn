package convert

import (
	"strings"

	"github.com/yuin/goldmark/v2/ast"
)

// lineBreak replaces line endings, which would split a code span over several
// lines and break the one-line structure of list items and table rows.
var lineBreak = strings.NewReplacer("\r\n", " ", "\n", " ")

// renderCodeSpan renders a code span in Cosense notation, which is the same
// `code` as in Markdown. Decorations are never applied to it.
// A backtick inside the code is written as is, because Cosense has no escape.
func (r *renderer) renderCodeSpan(span *ast.CodeSpan) string {
	return "`" + lineBreak.Replace(span.Value.Value(r.source)) + "`"
}
