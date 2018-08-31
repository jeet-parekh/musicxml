package musicxml

import (
	"encoding/xml"
	"strings"
)

// InstrumentAbbreviation is the structure for instrument-abbreviation MusicXML element.
type InstrumentAbbreviation struct {
	XMLName xml.Name `xml:"instrument-abbreviation"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *InstrumentAbbreviation) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("instrument-abbreviation", strings.TrimSpace(r.IValue), attributes, children)
}
