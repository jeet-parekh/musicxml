package musicxml

import (
	"encoding/xml"
	"strings"
)

// Release is the structure for release MusicXML element.
type Release struct {
	XMLName xml.Name `xml:"release"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Release) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("release", strings.TrimSpace(r.IValue), attributes, children)
}
