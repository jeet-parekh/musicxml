package mxml

import (
	"encoding/xml"
	"strings"
)

// NormalDot is the structure for normal-dot MusicXML element.
type NormalDot struct {
	XMLName xml.Name `xml:"normal-dot"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *NormalDot) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("normal-dot", strings.TrimSpace(r.IValue), attributes, children)
}
