package mxml

import (
	"encoding/xml"
	"strings"
)

// MidiBank is the structure for midi-bank MusicXML element.
type MidiBank struct {
	XMLName xml.Name `xml:"midi-bank"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MidiBank) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("midi-bank", strings.TrimSpace(r.IValue), attributes, children)
}
