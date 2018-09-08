package mxml

import (
	"encoding/xml"
	"strings"
)

// Pf is the structure for pf MusicXML element.
type Pf struct {
	XMLName xml.Name `xml:"pf"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Pf) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("pf", strings.TrimSpace(r.IValue), attributes, children)
}
