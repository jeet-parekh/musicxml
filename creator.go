package musicxml

import (
	"encoding/xml"
	"strings"
)

// Creator is the structure for creator MusicXML element.
type Creator struct {
	XMLName xml.Name `xml:"creator"`

	TypeAttr string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Creator) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["type"] = r.TypeAttr

	return newMXML("creator", strings.TrimSpace(r.IValue), attributes, children)
}
