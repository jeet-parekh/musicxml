package musicxml

import (
	"encoding/xml"
	"strings"
)

// MetronomeType is the structure for metronome-type MusicXML element.
type MetronomeType struct {
	XMLName xml.Name `xml:"metronome-type"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MetronomeType) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("metronome-type", strings.TrimSpace(r.IValue), attributes, children)
}
