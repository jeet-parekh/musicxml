package musicxml

import (
	"encoding/xml"
	"strings"
)

// Supports is the structure for supports MusicXML element.
type Supports struct {
	XMLName xml.Name `xml:"supports"`

	AttributeAttr string `xml:"attribute,attr,omitempty"`
	ElementAttr   string `xml:"element,attr,omitempty"`
	TypeAttr      string `xml:"type,attr,omitempty"`
	ValueAttr     string `xml:"value,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Supports) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["attribute"] = r.AttributeAttr
	attributes["element"] = r.ElementAttr
	attributes["type"] = r.TypeAttr
	attributes["value"] = r.ValueAttr

	return newMXML("supports", strings.TrimSpace(r.IValue), attributes, children)
}
