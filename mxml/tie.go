package mxml

import (
	"encoding/xml"
	"strings"
)

// Tie is the structure for tie MusicXML element.
type Tie struct {
	XMLName xml.Name `xml:"tie"`

	TimeOnlyAttr string `xml:"time-only,attr,omitempty"`
	TypeAttr     string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Tie) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["time-only"] = r.TimeOnlyAttr
	attributes["type"] = r.TypeAttr

	return newMXML("tie", strings.TrimSpace(r.IValue), attributes, children)
}
