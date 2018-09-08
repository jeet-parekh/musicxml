package mxml

import (
	"encoding/xml"
	"strings"
)

// TopMargin is the structure for top-margin MusicXML element.
type TopMargin struct {
	XMLName xml.Name `xml:"top-margin"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *TopMargin) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("top-margin", strings.TrimSpace(r.IValue), attributes, children)
}
