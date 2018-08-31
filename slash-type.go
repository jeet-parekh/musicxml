package musicxml

import (
	"encoding/xml"
	"strings"
)

// SlashType is the structure for slash-type MusicXML element.
type SlashType struct {
	XMLName xml.Name `xml:"slash-type"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *SlashType) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("slash-type", strings.TrimSpace(r.IValue), attributes, children)
}
