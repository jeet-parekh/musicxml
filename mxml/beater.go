package mxml

import (
	"encoding/xml"
	"strings"
)

// Beater is the structure for beater MusicXML element.
type Beater struct {
	XMLName xml.Name `xml:"beater"`

	TipAttr string `xml:"tip,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Beater) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["tip"] = r.TipAttr

	return newMXML("beater", strings.TrimSpace(r.IValue), attributes, children)
}
