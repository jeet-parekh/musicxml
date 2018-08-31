package musicxml

import (
	"encoding/xml"
	"strings"
)

// Ff is the structure for ff MusicXML element.
type Ff struct {
	XMLName xml.Name `xml:"ff"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Ff) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("ff", strings.TrimSpace(r.IValue), attributes, children)
}
