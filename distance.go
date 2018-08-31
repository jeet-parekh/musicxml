package musicxml

import (
	"encoding/xml"
	"strings"
)

// Distance is the structure for distance MusicXML element.
type Distance struct {
	XMLName xml.Name `xml:"distance"`

	TypeAttr string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Distance) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["type"] = r.TypeAttr

	return newMXML("distance", strings.TrimSpace(r.IValue), attributes, children)
}
