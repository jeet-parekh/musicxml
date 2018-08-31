package musicxml

import (
	"encoding/xml"
	"strings"
)

// CircularArrow is the structure for circular-arrow MusicXML element.
type CircularArrow struct {
	XMLName xml.Name `xml:"circular-arrow"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *CircularArrow) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("circular-arrow", strings.TrimSpace(r.IValue), attributes, children)
}
