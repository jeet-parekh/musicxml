package mxml

import (
	"encoding/xml"
	"strings"
)

// Rfz is the structure for rfz MusicXML element.
type Rfz struct {
	XMLName xml.Name `xml:"rfz"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Rfz) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("rfz", strings.TrimSpace(r.IValue), attributes, children)
}
