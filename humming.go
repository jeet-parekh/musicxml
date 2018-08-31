package musicxml

import (
	"encoding/xml"
	"strings"
)

// Humming is the structure for humming MusicXML element.
type Humming struct {
	XMLName xml.Name `xml:"humming"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Humming) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("humming", strings.TrimSpace(r.IValue), attributes, children)
}
