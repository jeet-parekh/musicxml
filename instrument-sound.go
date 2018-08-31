package musicxml

import (
	"encoding/xml"
	"strings"
)

// InstrumentSound is the structure for instrument-sound MusicXML element.
type InstrumentSound struct {
	XMLName xml.Name `xml:"instrument-sound"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *InstrumentSound) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("instrument-sound", strings.TrimSpace(r.IValue), attributes, children)
}
