package mxml

import (
	"encoding/xml"
	"strings"
)

// Sfpp is the structure for sfpp MusicXML element.
type Sfpp struct {
	XMLName xml.Name `xml:"sfpp"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Sfpp) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("sfpp", strings.TrimSpace(r.IValue), attributes, children)
}
