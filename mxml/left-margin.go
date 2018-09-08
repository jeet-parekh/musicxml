package mxml

import (
	"encoding/xml"
	"strings"
)

// LeftMargin is the structure for left-margin MusicXML element.
type LeftMargin struct {
	XMLName xml.Name `xml:"left-margin"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *LeftMargin) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("left-margin", strings.TrimSpace(r.IValue), attributes, children)
}
