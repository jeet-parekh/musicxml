package mxml

import (
	"encoding/xml"
	"strings"
)

// WorkTitle is the structure for work-title MusicXML element.
type WorkTitle struct {
	XMLName xml.Name `xml:"work-title"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *WorkTitle) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("work-title", strings.TrimSpace(r.IValue), attributes, children)
}
