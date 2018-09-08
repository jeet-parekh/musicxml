package mxml

import (
	"encoding/xml"
	"strings"
)

// Diatonic is the structure for diatonic MusicXML element.
type Diatonic struct {
	XMLName xml.Name `xml:"diatonic"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Diatonic) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("diatonic", strings.TrimSpace(r.IValue), attributes, children)
}
