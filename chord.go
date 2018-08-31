package musicxml

import (
	"encoding/xml"
	"strings"
)

// Chord is the structure for chord MusicXML element.
type Chord struct {
	XMLName xml.Name `xml:"chord"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Chord) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("chord", strings.TrimSpace(r.IValue), attributes, children)
}
