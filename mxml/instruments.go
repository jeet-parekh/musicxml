package mxml

import (
	"encoding/xml"
	"strings"
)

// Instruments is the structure for instruments MusicXML element.
type Instruments struct {
	XMLName xml.Name `xml:"instruments"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Instruments) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("instruments", strings.TrimSpace(r.IValue), attributes, children)
}
