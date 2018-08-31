package musicxml

import (
	"encoding/xml"
	"strings"
)

// Level is the structure for level MusicXML element.
type Level struct {
	XMLName xml.Name `xml:"level"`

	BracketAttr     string `xml:"bracket,attr,omitempty"`
	ParenthesesAttr string `xml:"parentheses,attr,omitempty"`
	ReferenceAttr   string `xml:"reference,attr,omitempty"`
	SizeAttr        string `xml:"size,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Level) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["bracket"] = r.BracketAttr
	attributes["parentheses"] = r.ParenthesesAttr
	attributes["reference"] = r.ReferenceAttr
	attributes["size"] = r.SizeAttr

	return newMXML("level", strings.TrimSpace(r.IValue), attributes, children)
}
