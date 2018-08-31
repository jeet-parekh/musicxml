package musicxml

import (
	"encoding/xml"
	"strings"
)

// OtherAppearance is the structure for other-appearance MusicXML element.
type OtherAppearance struct {
	XMLName xml.Name `xml:"other-appearance"`

	TypeAttr string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *OtherAppearance) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["type"] = r.TypeAttr

	return newMXML("other-appearance", strings.TrimSpace(r.IValue), attributes, children)
}
