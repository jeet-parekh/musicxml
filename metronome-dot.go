package musicxml

import (
	"encoding/xml"
	"strings"
)

// MetronomeDot is the structure for metronome-dot MusicXML element.
type MetronomeDot struct {
	XMLName xml.Name `xml:"metronome-dot"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MetronomeDot) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("metronome-dot", strings.TrimSpace(r.IValue), attributes, children)
}
