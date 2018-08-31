package musicxml

import (
	"encoding/xml"
	"strings"
)

// Title is the structure for title MusicXML element.
type Title struct {
	XMLName xml.Name `xml:"title"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Title) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("title", strings.TrimSpace(r.IValue), attributes, children)
}
