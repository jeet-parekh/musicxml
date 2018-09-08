package mxml

import (
	"encoding/xml"
	"strings"
)

// RightMargin is the structure for right-margin MusicXML element.
type RightMargin struct {
	XMLName xml.Name `xml:"right-margin"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *RightMargin) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("right-margin", strings.TrimSpace(r.IValue), attributes, children)
}
