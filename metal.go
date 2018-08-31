package musicxml

import (
	"encoding/xml"
	"strings"
)

// Metal is the structure for metal MusicXML element.
type Metal struct {
	XMLName xml.Name `xml:"metal"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Metal) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("metal", strings.TrimSpace(r.IValue), attributes, children)
}
