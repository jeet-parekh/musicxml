package musicxml

import (
	"encoding/xml"
	"strings"
)

// Accidental is the structure for accidental MusicXML element.
type Accidental struct {
	XMLName xml.Name `xml:"accidental"`

	BracketAttr     string `xml:"bracket,attr,omitempty"`
	CautionaryAttr  string `xml:"cautionary,attr,omitempty"`
	ColorAttr       string `xml:"color,attr,omitempty"`
	DefaultXAttr    string `xml:"default-x,attr,omitempty"`
	DefaultYAttr    string `xml:"default-y,attr,omitempty"`
	EditorialAttr   string `xml:"editorial,attr,omitempty"`
	FontFamilyAttr  string `xml:"font-family,attr,omitempty"`
	FontSizeAttr    string `xml:"font-size,attr,omitempty"`
	FontStyleAttr   string `xml:"font-style,attr,omitempty"`
	FontWeightAttr  string `xml:"font-weight,attr,omitempty"`
	ParenthesesAttr string `xml:"parentheses,attr,omitempty"`
	RelativeXAttr   string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr   string `xml:"relative-y,attr,omitempty"`
	SizeAttr        string `xml:"size,attr,omitempty"`
	SmuflAttr       string `xml:"smufl,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Accidental) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["bracket"] = r.BracketAttr
	attributes["cautionary"] = r.CautionaryAttr
	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["editorial"] = r.EditorialAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["parentheses"] = r.ParenthesesAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["size"] = r.SizeAttr
	attributes["smufl"] = r.SmuflAttr

	return newMXML("accidental", strings.TrimSpace(r.IValue), attributes, children)
}
