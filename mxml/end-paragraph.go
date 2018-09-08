package mxml

import (
	"encoding/xml"
	"strings"
)

// EndParagraph is the structure for end-paragraph MusicXML element.
type EndParagraph struct {
	XMLName xml.Name `xml:"end-paragraph"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *EndParagraph) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("end-paragraph", strings.TrimSpace(r.IValue), attributes, children)
}
