package mxml

import (
	"encoding/xml"
	"strings"
)

// KeyAccidental is the structure for key-accidental MusicXML element.
type KeyAccidental struct {
	XMLName xml.Name `xml:"key-accidental"`

	SmuflAttr string `xml:"smufl,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *KeyAccidental) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["smufl"] = r.SmuflAttr

	return newMXML("key-accidental", strings.TrimSpace(r.IValue), attributes, children)
}
