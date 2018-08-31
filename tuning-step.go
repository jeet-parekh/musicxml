package musicxml

import (
	"encoding/xml"
	"strings"
)

// TuningStep is the structure for tuning-step MusicXML element.
type TuningStep struct {
	XMLName xml.Name `xml:"tuning-step"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *TuningStep) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("tuning-step", strings.TrimSpace(r.IValue), attributes, children)
}
