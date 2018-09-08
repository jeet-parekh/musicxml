package mxml

import (
	"encoding/xml"
	"strings"
)

// Software is the structure for software MusicXML element.
type Software struct {
	XMLName xml.Name `xml:"software"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Software) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("software", strings.TrimSpace(r.IValue), attributes, children)
}
