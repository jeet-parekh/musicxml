package musicxml

import (
	"encoding/xml"
	"strings"
)

// Fff is the structure for fff MusicXML element.
type Fff struct {
	XMLName xml.Name `xml:"fff"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Fff) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("fff", strings.TrimSpace(r.IValue), attributes, children)
}
