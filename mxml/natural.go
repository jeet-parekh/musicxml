package mxml

import (
	"encoding/xml"
	"strings"
)

// Natural is the structure for natural MusicXML element.
type Natural struct {
	XMLName xml.Name `xml:"natural"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Natural) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("natural", strings.TrimSpace(r.IValue), attributes, children)
}
