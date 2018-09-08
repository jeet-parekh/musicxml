package mxml

import (
	"encoding/xml"
	"strings"
)

// Offset is the structure for offset MusicXML element.
type Offset struct {
	XMLName xml.Name `xml:"offset"`

	SoundAttr string `xml:"sound,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Offset) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["sound"] = r.SoundAttr

	return newMXML("offset", strings.TrimSpace(r.IValue), attributes, children)
}
