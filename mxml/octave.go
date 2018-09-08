package mxml

import (
	"encoding/xml"
	"strings"
)

// Octave is the structure for octave MusicXML element.
type Octave struct {
	XMLName xml.Name `xml:"octave"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Octave) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("octave", strings.TrimSpace(r.IValue), attributes, children)
}
