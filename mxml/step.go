package mxml

import (
	"encoding/xml"
	"strings"
)

// Step is the structure for step MusicXML element.
type Step struct {
	XMLName xml.Name `xml:"step"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Step) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("step", strings.TrimSpace(r.IValue), attributes, children)
}
