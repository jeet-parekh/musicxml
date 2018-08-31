package musicxml

import (
	"encoding/xml"
	"strings"
)

// MidiChannel is the structure for midi-channel MusicXML element.
type MidiChannel struct {
	XMLName xml.Name `xml:"midi-channel"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MidiChannel) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("midi-channel", strings.TrimSpace(r.IValue), attributes, children)
}
