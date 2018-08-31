package musicxml

import (
	"encoding/xml"
	"strings"
)

// AccidentalMark is the structure for accidental-mark MusicXML element.
type AccidentalMark struct {
	XMLName xml.Name `xml:"accidental-mark"`

	BracketAttr     string `xml:"bracket,attr,omitempty"`
	ColorAttr       string `xml:"color,attr,omitempty"`
	DefaultXAttr    string `xml:"default-x,attr,omitempty"`
	DefaultYAttr    string `xml:"default-y,attr,omitempty"`
	FontFamilyAttr  string `xml:"font-family,attr,omitempty"`
	FontSizeAttr    string `xml:"font-size,attr,omitempty"`
	FontStyleAttr   string `xml:"font-style,attr,omitempty"`
	FontWeightAttr  string `xml:"font-weight,attr,omitempty"`
	IdAttr          string `xml:"id,attr,omitempty"`
	ParenthesesAttr string `xml:"parentheses,attr,omitempty"`
	PlacementAttr   string `xml:"placement,attr,omitempty"`
	RelativeXAttr   string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr   string `xml:"relative-y,attr,omitempty"`
	SizeAttr        string `xml:"size,attr,omitempty"`
	SmuflAttr       string `xml:"smufl,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *AccidentalMark) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["bracket"] = r.BracketAttr
	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["id"] = r.IdAttr
	attributes["parentheses"] = r.ParenthesesAttr
	attributes["placement"] = r.PlacementAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["size"] = r.SizeAttr
	attributes["smufl"] = r.SmuflAttr

	return newMXML("accidental-mark", strings.TrimSpace(r.IValue), attributes, children)
}
