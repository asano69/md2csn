package convert

import "strings"

// style is a set of decorations. Inline rendering passes the style inherited
// from the enclosing nodes down to the text, so nested decorations end up as
// one bracket per text.
type style uint8

const styleNone style = 0

const (
	styleBold style = 1 << iota
	styleItalic
	styleStrike
	// stylePlain drops every decoration below it, as in headings and table cells.
	stylePlain
)

// add returns s with d added. A plain style stays plain.
func (s style) add(d style) style {
	if s&stylePlain != 0 {
		return s
	}
	return s | d
}

// marks returns the Cosense symbols of s in a fixed order: "*", "/", "-".
func (s style) marks() string {
	var sb strings.Builder
	if s&styleBold != 0 {
		sb.WriteByte('*')
	}
	if s&styleItalic != 0 {
		sb.WriteByte('/')
	}
	if s&styleStrike != 0 {
		sb.WriteByte('-')
	}
	return sb.String()
}

// wrap puts text in one bracket with the marks of s, such as "[*/ text]".
// Whitespace around text stays outside the bracket. Text is returned as is
// when s has no marks or text is blank.
func (s style) wrap(text string) string {
	marks := s.marks()
	core := strings.TrimSpace(text)
	if marks == "" || core == "" {
		return text
	}
	start := strings.Index(text, core)
	return text[:start] + "[" + marks + " " + core + "]" + text[start+len(core):]
}
