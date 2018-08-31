package musicxml

import (
	"encoding/xml"
	"strings"
)

// Line is the structure for line MusicXML element.
type Line struct {
	XMLName xml.Name `xml:"line"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Line) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("line", strings.TrimSpace(r.IValue), attributes, children)
}
