package musicxml

import (
	"encoding/xml"
	"strings"
)

// MetronomeRelation is the structure for metronome-relation MusicXML element.
type MetronomeRelation struct {
	XMLName xml.Name `xml:"metronome-relation"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MetronomeRelation) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("metronome-relation", strings.TrimSpace(r.IValue), attributes, children)
}
