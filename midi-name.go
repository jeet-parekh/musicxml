package musicxml

import (
	"encoding/xml"
	"strings"
)

// MidiName is the structure for midi-name MusicXML element.
type MidiName struct {
	XMLName xml.Name `xml:"midi-name"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MidiName) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("midi-name", strings.TrimSpace(r.IValue), attributes, children)
}
