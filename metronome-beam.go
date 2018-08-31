package musicxml

import (
	"encoding/xml"
	"strings"
)

// MetronomeBeam is the structure for metronome-beam MusicXML element.
type MetronomeBeam struct {
	XMLName xml.Name `xml:"metronome-beam"`

	NumberAttr string `xml:"number,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MetronomeBeam) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["number"] = r.NumberAttr

	return newMXML("metronome-beam", strings.TrimSpace(r.IValue), attributes, children)
}
