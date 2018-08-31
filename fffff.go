package musicxml

import (
	"encoding/xml"
	"strings"
)

// Fffff is the structure for fffff MusicXML element.
type Fffff struct {
	XMLName xml.Name `xml:"fffff"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Fffff) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("fffff", strings.TrimSpace(r.IValue), attributes, children)
}
