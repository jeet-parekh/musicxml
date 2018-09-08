package mxml

import (
	"encoding/xml"
	"strings"
)

// Sf is the structure for sf MusicXML element.
type Sf struct {
	XMLName xml.Name `xml:"sf"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Sf) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("sf", strings.TrimSpace(r.IValue), attributes, children)
}
