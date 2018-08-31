package musicxml

import (
	"encoding/xml"
	"strings"
)

// Artificial is the structure for artificial MusicXML element.
type Artificial struct {
	XMLName xml.Name `xml:"artificial"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Artificial) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("artificial", strings.TrimSpace(r.IValue), attributes, children)
}
