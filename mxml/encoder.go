package mxml

import (
	"encoding/xml"
	"strings"
)

// Encoder is the structure for encoder MusicXML element.
type Encoder struct {
	XMLName xml.Name `xml:"encoder"`

	TypeAttr string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Encoder) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["type"] = r.TypeAttr

	return newMXML("encoder", strings.TrimSpace(r.IValue), attributes, children)
}
