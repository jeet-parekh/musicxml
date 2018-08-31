package musicxml

import (
	"encoding/xml"
	"strings"
)

// PreBend is the structure for pre-bend MusicXML element.
type PreBend struct {
	XMLName xml.Name `xml:"pre-bend"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *PreBend) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("pre-bend", strings.TrimSpace(r.IValue), attributes, children)
}
