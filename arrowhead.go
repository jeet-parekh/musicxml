package musicxml

import (
	"encoding/xml"
	"strings"
)

// Arrowhead is the structure for arrowhead MusicXML element.
type Arrowhead struct {
	XMLName xml.Name `xml:"arrowhead"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Arrowhead) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("arrowhead", strings.TrimSpace(r.IValue), attributes, children)
}
