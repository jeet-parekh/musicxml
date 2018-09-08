package mxml

import (
	"encoding/xml"
	"strings"
)

// Membrane is the structure for membrane MusicXML element.
type Membrane struct {
	XMLName xml.Name `xml:"membrane"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Membrane) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("membrane", strings.TrimSpace(r.IValue), attributes, children)
}
