package mxml

import (
	"encoding/xml"
	"strings"
)

// Pppp is the structure for pppp MusicXML element.
type Pppp struct {
	XMLName xml.Name `xml:"pppp"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Pppp) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("pppp", strings.TrimSpace(r.IValue), attributes, children)
}
