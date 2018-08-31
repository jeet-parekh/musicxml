package musicxml

import (
	"encoding/xml"
	"strings"
)

// MetronomeTied is the structure for metronome-tied MusicXML element.
type MetronomeTied struct {
	XMLName xml.Name `xml:"metronome-tied"`

	TypeAttr string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MetronomeTied) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["type"] = r.TypeAttr

	return newMXML("metronome-tied", strings.TrimSpace(r.IValue), attributes, children)
}
