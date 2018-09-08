package mxml

import (
	"encoding/xml"
	"strings"
)

// Relation is the structure for relation MusicXML element.
type Relation struct {
	XMLName xml.Name `xml:"relation"`

	TypeAttr string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Relation) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["type"] = r.TypeAttr

	return newMXML("relation", strings.TrimSpace(r.IValue), attributes, children)
}
