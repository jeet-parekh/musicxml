package mxml

import (
	"encoding/xml"
	"strings"
)

// PageWidth is the structure for page-width MusicXML element.
type PageWidth struct {
	XMLName xml.Name `xml:"page-width"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *PageWidth) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("page-width", strings.TrimSpace(r.IValue), attributes, children)
}
