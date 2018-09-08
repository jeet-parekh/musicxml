package mxml

import (
	"encoding/xml"
	"strings"
)

// Sffz is the structure for sffz MusicXML element.
type Sffz struct {
	XMLName xml.Name `xml:"sffz"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Sffz) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("sffz", strings.TrimSpace(r.IValue), attributes, children)
}
