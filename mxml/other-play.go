package mxml

import (
	"encoding/xml"
	"strings"
)

// OtherPlay is the structure for other-play MusicXML element.
type OtherPlay struct {
	XMLName xml.Name `xml:"other-play"`

	TypeAttr string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *OtherPlay) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["type"] = r.TypeAttr

	return newMXML("other-play", strings.TrimSpace(r.IValue), attributes, children)
}
