package musicxml

import (
	"encoding/xml"
	"strings"
)

// Ffffff is the structure for ffffff MusicXML element.
type Ffffff struct {
	XMLName xml.Name `xml:"ffffff"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Ffffff) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("ffffff", strings.TrimSpace(r.IValue), attributes, children)
}
