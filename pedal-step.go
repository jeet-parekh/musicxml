package musicxml

import (
	"encoding/xml"
	"strings"
)

// PedalStep is the structure for pedal-step MusicXML element.
type PedalStep struct {
	XMLName xml.Name `xml:"pedal-step"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *PedalStep) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("pedal-step", strings.TrimSpace(r.IValue), attributes, children)
}
