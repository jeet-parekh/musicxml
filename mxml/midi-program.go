package mxml

import (
	"encoding/xml"
	"strings"
)

// MidiProgram is the structure for midi-program MusicXML element.
type MidiProgram struct {
	XMLName xml.Name `xml:"midi-program"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MidiProgram) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("midi-program", strings.TrimSpace(r.IValue), attributes, children)
}
