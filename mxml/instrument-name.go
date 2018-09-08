package mxml

import (
	"encoding/xml"
	"strings"
)

// InstrumentName is the structure for instrument-name MusicXML element.
type InstrumentName struct {
	XMLName xml.Name `xml:"instrument-name"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *InstrumentName) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("instrument-name", strings.TrimSpace(r.IValue), attributes, children)
}
