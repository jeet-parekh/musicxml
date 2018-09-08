package mxml

import (
	"encoding/xml"
	"strings"
)

// MetronomeArrows is the structure for metronome-arrows MusicXML element.
type MetronomeArrows struct {
	XMLName xml.Name `xml:"metronome-arrows"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MetronomeArrows) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("metronome-arrows", strings.TrimSpace(r.IValue), attributes, children)
}
