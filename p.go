package musicxml

import (
	"encoding/xml"
	"strings"
)

// P is the structure for p MusicXML element.
type P struct {
	XMLName xml.Name `xml:"p"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *P) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("p", strings.TrimSpace(r.IValue), attributes, children)
}
