package mxml

import (
	"encoding/xml"
	"strings"
)

// Ppp is the structure for ppp MusicXML element.
type Ppp struct {
	XMLName xml.Name `xml:"ppp"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Ppp) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("ppp", strings.TrimSpace(r.IValue), attributes, children)
}
