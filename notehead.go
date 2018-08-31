package musicxml

import (
	"encoding/xml"
	"strings"
)

// Notehead is the structure for notehead MusicXML element.
type Notehead struct {
	XMLName xml.Name `xml:"notehead"`

	ColorAttr       string `xml:"color,attr,omitempty"`
	FilledAttr      string `xml:"filled,attr,omitempty"`
	FontFamilyAttr  string `xml:"font-family,attr,omitempty"`
	FontSizeAttr    string `xml:"font-size,attr,omitempty"`
	FontStyleAttr   string `xml:"font-style,attr,omitempty"`
	FontWeightAttr  string `xml:"font-weight,attr,omitempty"`
	ParenthesesAttr string `xml:"parentheses,attr,omitempty"`
	SmuflAttr       string `xml:"smufl,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Notehead) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["filled"] = r.FilledAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["parentheses"] = r.ParenthesesAttr
	attributes["smufl"] = r.SmuflAttr

	return newMXML("notehead", strings.TrimSpace(r.IValue), attributes, children)
}
