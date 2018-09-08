package mxml

import (
	"encoding/xml"
	"strings"
)

// Glyph is the structure for glyph MusicXML element.
type Glyph struct {
	XMLName xml.Name `xml:"glyph"`

	TypeAttr string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Glyph) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["type"] = r.TypeAttr

	return newMXML("glyph", strings.TrimSpace(r.IValue), attributes, children)
}
