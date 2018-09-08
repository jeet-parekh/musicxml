package mxml

import (
	"encoding/xml"
	"strings"
)

// Fifths is the structure for fifths MusicXML element.
type Fifths struct {
	XMLName xml.Name `xml:"fifths"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Fifths) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("fifths", strings.TrimSpace(r.IValue), attributes, children)
}
