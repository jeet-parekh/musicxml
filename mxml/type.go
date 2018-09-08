package mxml

import (
	"encoding/xml"
	"strings"
)

// Type is the structure for type MusicXML element.
type Type struct {
	XMLName xml.Name `xml:"type"`

	SizeAttr string `xml:"size,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Type) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["size"] = r.SizeAttr

	return newMXML("type", strings.TrimSpace(r.IValue), attributes, children)
}
