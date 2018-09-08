package mxml

import (
	"encoding/xml"
	"strings"
)

// Voice is the structure for voice MusicXML element.
type Voice struct {
	XMLName xml.Name `xml:"voice"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Voice) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("voice", strings.TrimSpace(r.IValue), attributes, children)
}
