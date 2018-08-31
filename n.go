package musicxml

import (
	"encoding/xml"
	"strings"
)

// N is the structure for n MusicXML element.
type N struct {
	XMLName xml.Name `xml:"n"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *N) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("n", strings.TrimSpace(r.IValue), attributes, children)
}
