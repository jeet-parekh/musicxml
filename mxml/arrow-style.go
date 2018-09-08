package mxml

import (
	"encoding/xml"
	"strings"
)

// ArrowStyle is the structure for arrow-style MusicXML element.
type ArrowStyle struct {
	XMLName xml.Name `xml:"arrow-style"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *ArrowStyle) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("arrow-style", strings.TrimSpace(r.IValue), attributes, children)
}
