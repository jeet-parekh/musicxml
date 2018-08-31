package musicxml

import (
	"encoding/xml"
	"strings"
)

// Pan is the structure for pan MusicXML element.
type Pan struct {
	XMLName xml.Name `xml:"pan"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Pan) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("pan", strings.TrimSpace(r.IValue), attributes, children)
}
