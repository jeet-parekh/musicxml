package mxml

import (
	"encoding/xml"
	"strings"
)

// Mode is the structure for mode MusicXML element.
type Mode struct {
	XMLName xml.Name `xml:"mode"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Mode) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("mode", strings.TrimSpace(r.IValue), attributes, children)
}
