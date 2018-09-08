package mxml

import (
	"encoding/xml"
	"strings"
)

// Pp is the structure for pp MusicXML element.
type Pp struct {
	XMLName xml.Name `xml:"pp"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Pp) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("pp", strings.TrimSpace(r.IValue), attributes, children)
}
