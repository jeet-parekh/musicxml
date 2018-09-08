package mxml

import (
	"encoding/xml"
	"strings"
)

// Ipa is the structure for ipa MusicXML element.
type Ipa struct {
	XMLName xml.Name `xml:"ipa"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Ipa) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("ipa", strings.TrimSpace(r.IValue), attributes, children)
}
