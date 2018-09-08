package mxml

import (
	"encoding/xml"
	"strings"
)

// Mf is the structure for mf MusicXML element.
type Mf struct {
	XMLName xml.Name `xml:"mf"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Mf) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("mf", strings.TrimSpace(r.IValue), attributes, children)
}
