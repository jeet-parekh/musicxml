package musicxml

import (
	"encoding/xml"
	"strings"
)

// Rf is the structure for rf MusicXML element.
type Rf struct {
	XMLName xml.Name `xml:"rf"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Rf) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("rf", strings.TrimSpace(r.IValue), attributes, children)
}
