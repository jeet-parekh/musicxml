package mxml

import (
	"encoding/xml"
	"strings"
)

// DisplayStep is the structure for display-step MusicXML element.
type DisplayStep struct {
	XMLName xml.Name `xml:"display-step"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *DisplayStep) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("display-step", strings.TrimSpace(r.IValue), attributes, children)
}
