package musicxml

import (
	"encoding/xml"
	"strings"
)

// Sfz is the structure for sfz MusicXML element.
type Sfz struct {
	XMLName xml.Name `xml:"sfz"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Sfz) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("sfz", strings.TrimSpace(r.IValue), attributes, children)
}
