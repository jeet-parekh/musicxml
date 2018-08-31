package musicxml

import (
	"encoding/xml"
	"strings"
)

// Rights is the structure for rights MusicXML element.
type Rights struct {
	XMLName xml.Name `xml:"rights"`

	TypeAttr string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Rights) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["type"] = r.TypeAttr

	return newMXML("rights", strings.TrimSpace(r.IValue), attributes, children)
}
