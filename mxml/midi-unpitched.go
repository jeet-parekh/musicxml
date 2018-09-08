package mxml

import (
	"encoding/xml"
	"strings"
)

// MidiUnpitched is the structure for midi-unpitched MusicXML element.
type MidiUnpitched struct {
	XMLName xml.Name `xml:"midi-unpitched"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MidiUnpitched) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("midi-unpitched", strings.TrimSpace(r.IValue), attributes, children)
}
