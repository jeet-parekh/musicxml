package musicxml

import (
	"encoding/xml"
	"strings"
)

// Source is the structure for source MusicXML element.
type Source struct {
	XMLName xml.Name `xml:"source"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Source) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("source", strings.TrimSpace(r.IValue), attributes, children)
}
