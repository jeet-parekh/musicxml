package mxml

import (
	"encoding/xml"
	"strings"
)

// BendAlter is the structure for bend-alter MusicXML element.
type BendAlter struct {
	XMLName xml.Name `xml:"bend-alter"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *BendAlter) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("bend-alter", strings.TrimSpace(r.IValue), attributes, children)
}
