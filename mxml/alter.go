package mxml

import (
	"encoding/xml"
	"strings"
)

// Alter is the structure for alter MusicXML element.
type Alter struct {
	XMLName xml.Name `xml:"alter"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Alter) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("alter", strings.TrimSpace(r.IValue), attributes, children)
}
