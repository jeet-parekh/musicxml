package mxml

import (
	"encoding/xml"
	"strings"
)

// Fz is the structure for fz MusicXML element.
type Fz struct {
	XMLName xml.Name `xml:"fz"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Fz) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("fz", strings.TrimSpace(r.IValue), attributes, children)
}
