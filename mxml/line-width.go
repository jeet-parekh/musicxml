package mxml

import (
	"encoding/xml"
	"strings"
)

// LineWidth is the structure for line-width MusicXML element.
type LineWidth struct {
	XMLName xml.Name `xml:"line-width"`

	TypeAttr string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *LineWidth) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["type"] = r.TypeAttr

	return newMXML("line-width", strings.TrimSpace(r.IValue), attributes, children)
}
