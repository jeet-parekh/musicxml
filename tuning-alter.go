package musicxml

import (
	"encoding/xml"
	"strings"
)

// TuningAlter is the structure for tuning-alter MusicXML element.
type TuningAlter struct {
	XMLName xml.Name `xml:"tuning-alter"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *TuningAlter) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("tuning-alter", strings.TrimSpace(r.IValue), attributes, children)
}
