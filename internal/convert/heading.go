package convert

import (
	"slices"

	"github.com/yuin/goldmark/v2/ast"
)

// headingStars maps each Markdown heading level used in the document to
// its number of asterisks in Cosense notation.
//
// A single asterisk is reserved for bold, so the deepest level used gets two
// and each shallower level gets one more. Unused levels are skipped.
func headingStars(doc ast.Node) map[int]int {
	levels := collectHeadingLevels(doc, nil)
	slices.Sort(levels)
	levels = slices.Compact(levels)

	stars := make(map[int]int, len(levels))
	for i, level := range levels {
		stars[level] = len(levels) - i + 1
	}
	return stars
}

// collectHeadingLevels appends the level of every heading under n to levels.
func collectHeadingLevels(n ast.Node, levels []int) []int {
	if h, ok := n.(*ast.Heading); ok {
		levels = append(levels, h.Level)
	}
	for c := range children(n) {
		levels = collectHeadingLevels(c, levels)
	}
	return levels
}
