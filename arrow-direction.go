package musicxml

import (
	"encoding/xml"
	"strings"
)

// ArrowDirection is the structure for arrow-direction MusicXML element.
type ArrowDirection struct {
	XMLName xml.Name `xml:"arrow-direction"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *ArrowDirection) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("arrow-direction", strings.TrimSpace(r.IValue), attributes, children)
}
