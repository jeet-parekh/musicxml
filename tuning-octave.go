package musicxml

import (
	"encoding/xml"
	"strings"
)

// TuningOctave is the structure for tuning-octave MusicXML element.
type TuningOctave struct {
	XMLName xml.Name `xml:"tuning-octave"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *TuningOctave) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("tuning-octave", strings.TrimSpace(r.IValue), attributes, children)
}
