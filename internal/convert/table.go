package convert

import (
	"strings"

	"github.com/yuin/goldmark/v2/ast"
	extast "github.com/yuin/goldmark/v2/extension/ast"
)

// cellSpace replaces characters that would break the row layout.
var cellSpace = strings.NewReplacer("\r\n", " ", "\n", " ", "\t", " ")

// renderTable renders a GFM table in Cosense notation: a "table:_" line
// followed by one tab-indented line per row, with the cells separated by tabs.
// The header is written as a normal row. The delimiter row and the column
// alignments are dropped.
func (r *renderer) renderTable(table *extast.Table) string {
	lines := []string{"table:_"}
	for _, row := range tableRows(table, nil) {
		lines = append(lines, "\t"+r.renderRow(row))
	}
	return strings.Join(lines, "\n")
}

// tableRows appends every row under n to rows, in document order.
// A row is a node whose first child is a cell, so this works whether the
// header holds a row node or the cells directly.
func tableRows(n ast.Node, rows []ast.Node) []ast.Node {
	for c := range children(n) {
		if _, ok := c.FirstChild().(*extast.TableCell); ok {
			rows = append(rows, c)
		} else {
			rows = tableRows(c, rows)
		}
	}
	return rows
}

// renderRow renders the cells of one row separated by tabs.
// Cell content is plain inline text, so only [[wikilinks]] are converted
// and decorations are dropped.
func (r *renderer) renderRow(row ast.Node) string {
	var cells []string
	for c := range children(row) {
		cells = append(cells, cellSpace.Replace(r.renderInlines(c, stylePlain)))
	}
	return strings.Join(cells, "\t")
}
