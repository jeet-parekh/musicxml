package musicxml

import (
	"encoding/xml"
	"strings"
)

// Effect is the structure for effect MusicXML element.
type Effect struct {
	XMLName xml.Name `xml:"effect"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Effect) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("effect", strings.TrimSpace(r.IValue), attributes, children)
}
