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
// The skip heading, if any, is not counted because it is not rendered.
func headingStars(doc, skip ast.Node) map[int]int {
	levels := collectHeadingLevels(doc, skip, nil)
	slices.Sort(levels)
	levels = slices.Compact(levels)

	stars := make(map[int]int, len(levels))
	for i, level := range levels {
		stars[level] = len(levels) - i + 1
	}
	return stars
}

// collectHeadingLevels appends the level of every heading under n to levels,
// except for the skip heading.
func collectHeadingLevels(n, skip ast.Node, levels []int) []int {
	if h, ok := n.(*ast.Heading); ok && n != skip {
		levels = append(levels, h.Level)
	}
	for c := range children(n) {
		levels = collectHeadingLevels(c, skip, levels)
	}
	return levels
}
