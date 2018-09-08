package mxml

import (
	"encoding/xml"
	"strings"
)

// Tenths is the structure for tenths MusicXML element.
type Tenths struct {
	XMLName xml.Name `xml:"tenths"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Tenths) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("tenths", strings.TrimSpace(r.IValue), attributes, children)
}
