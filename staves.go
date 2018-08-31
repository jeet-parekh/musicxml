package musicxml

import (
	"encoding/xml"
	"strings"
)

// Staves is the structure for staves MusicXML element.
type Staves struct {
	XMLName xml.Name `xml:"staves"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Staves) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("staves", strings.TrimSpace(r.IValue), attributes, children)
}
