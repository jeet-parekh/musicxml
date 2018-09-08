package mxml

import (
	"encoding/xml"
	"strings"
)

// ActualNotes is the structure for actual-notes MusicXML element.
type ActualNotes struct {
	XMLName xml.Name `xml:"actual-notes"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *ActualNotes) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("actual-notes", strings.TrimSpace(r.IValue), attributes, children)
}
