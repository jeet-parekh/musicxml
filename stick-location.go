package musicxml

import (
	"encoding/xml"
	"strings"
)

// StickLocation is the structure for stick-location MusicXML element.
type StickLocation struct {
	XMLName xml.Name `xml:"stick-location"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *StickLocation) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("stick-location", strings.TrimSpace(r.IValue), attributes, children)
}
