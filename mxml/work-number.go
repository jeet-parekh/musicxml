package mxml

import (
	"encoding/xml"
	"strings"
)

// WorkNumber is the structure for work-number MusicXML element.
type WorkNumber struct {
	XMLName xml.Name `xml:"work-number"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *WorkNumber) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("work-number", strings.TrimSpace(r.IValue), attributes, children)
}
