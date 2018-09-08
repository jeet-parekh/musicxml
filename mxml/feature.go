package mxml

import (
	"encoding/xml"
	"strings"
)

// Feature is the structure for feature MusicXML element.
type Feature struct {
	XMLName xml.Name `xml:"feature"`

	TypeAttr string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Feature) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["type"] = r.TypeAttr

	return newMXML("feature", strings.TrimSpace(r.IValue), attributes, children)
}
