package mxml

import (
	"encoding/xml"
	"strings"
)

// NormalNotes is the structure for normal-notes MusicXML element.
type NormalNotes struct {
	XMLName xml.Name `xml:"normal-notes"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *NormalNotes) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("normal-notes", strings.TrimSpace(r.IValue), attributes, children)
}
