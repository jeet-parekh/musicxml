package musicxml

import (
	"encoding/xml"
	"strings"
)

// Duration is the structure for duration MusicXML element.
type Duration struct {
	XMLName xml.Name `xml:"duration"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Duration) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("duration", strings.TrimSpace(r.IValue), attributes, children)
}
