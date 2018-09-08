package mxml

import (
	"encoding/xml"
	"strings"
)

// PedalAlter is the structure for pedal-alter MusicXML element.
type PedalAlter struct {
	XMLName xml.Name `xml:"pedal-alter"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *PedalAlter) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("pedal-alter", strings.TrimSpace(r.IValue), attributes, children)
}
