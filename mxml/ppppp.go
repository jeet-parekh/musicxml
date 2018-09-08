package mxml

import (
	"encoding/xml"
	"strings"
)

// Ppppp is the structure for ppppp MusicXML element.
type Ppppp struct {
	XMLName xml.Name `xml:"ppppp"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Ppppp) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("ppppp", strings.TrimSpace(r.IValue), attributes, children)
}
