package musicxml

import (
	"encoding/xml"
	"strings"
)

// StickMaterial is the structure for stick-material MusicXML element.
type StickMaterial struct {
	XMLName xml.Name `xml:"stick-material"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *StickMaterial) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("stick-material", strings.TrimSpace(r.IValue), attributes, children)
}
