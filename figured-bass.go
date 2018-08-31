package musicxml

import (
	"encoding/xml"
	"strings"
)

// FiguredBass is the structure for figured-bass MusicXML element.
type FiguredBass struct {
	XMLName xml.Name `xml:"figured-bass"`

	ColorAttr        string `xml:"color,attr,omitempty"`
	DefaultXAttr     string `xml:"default-x,attr,omitempty"`
	DefaultYAttr     string `xml:"default-y,attr,omitempty"`
	FontFamilyAttr   string `xml:"font-family,attr,omitempty"`
	FontSizeAttr     string `xml:"font-size,attr,omitempty"`
	FontStyleAttr    string `xml:"font-style,attr,omitempty"`
	FontWeightAttr   string `xml:"font-weight,attr,omitempty"`
	IdAttr           string `xml:"id,attr,omitempty"`
	ParenthesesAttr  string `xml:"parentheses,attr,omitempty"`
	PrintDotAttr     string `xml:"print-dot,attr,omitempty"`
	PrintLyricAttr   string `xml:"print-lyric,attr,omitempty"`
	PrintObjectAttr  string `xml:"print-object,attr,omitempty"`
	PrintSpacingAttr string `xml:"print-spacing,attr,omitempty"`
	RelativeXAttr    string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr    string `xml:"relative-y,attr,omitempty"`

	Duration []Duration `xml:"duration,omitempty"`
	Figure   []Figure   `xml:"figure,omitempty"`
	Footnote []Footnote `xml:"footnote,omitempty"`
	Level    []Level    `xml:"level,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *FiguredBass) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["id"] = r.IdAttr
	attributes["parentheses"] = r.ParenthesesAttr
	attributes["print-dot"] = r.PrintDotAttr
	attributes["print-lyric"] = r.PrintLyricAttr
	attributes["print-object"] = r.PrintObjectAttr
	attributes["print-spacing"] = r.PrintSpacingAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr

	children["duration"] = make([]*MXML, len(r.Duration))
	for i, c := range r.Duration {
		children["duration"][i] = c.ToMXML()
	}

	children["figure"] = make([]*MXML, len(r.Figure))
	for i, c := range r.Figure {
		children["figure"][i] = c.ToMXML()
	}

	children["footnote"] = make([]*MXML, len(r.Footnote))
	for i, c := range r.Footnote {
		children["footnote"][i] = c.ToMXML()
	}

	children["level"] = make([]*MXML, len(r.Level))
	for i, c := range r.Level {
		children["level"][i] = c.ToMXML()
	}

	return newMXML("figured-bass", strings.TrimSpace(r.IValue), attributes, children)
}
