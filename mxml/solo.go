package mxml

import (
	"encoding/xml"
	"strings"
)

// Solo is the structure for solo MusicXML element.
type Solo struct {
	XMLName xml.Name `xml:"solo"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Solo) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("solo", strings.TrimSpace(r.IValue), attributes, children)
}
