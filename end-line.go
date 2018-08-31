package musicxml

import (
	"encoding/xml"
	"strings"
)

// EndLine is the structure for end-line MusicXML element.
type EndLine struct {
	XMLName xml.Name `xml:"end-line"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *EndLine) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("end-line", strings.TrimSpace(r.IValue), attributes, children)
}
