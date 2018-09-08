package mxml

import (
	"encoding/xml"
	"strings"
)

// Chromatic is the structure for chromatic MusicXML element.
type Chromatic struct {
	XMLName xml.Name `xml:"chromatic"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Chromatic) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("chromatic", strings.TrimSpace(r.IValue), attributes, children)
}
