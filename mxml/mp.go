package mxml

import (
	"encoding/xml"
	"strings"
)

// Mp is the structure for mp MusicXML element.
type Mp struct {
	XMLName xml.Name `xml:"mp"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Mp) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("mp", strings.TrimSpace(r.IValue), attributes, children)
}
