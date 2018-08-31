package musicxml

import (
	"encoding/xml"
	"strings"
)

// SlashDot is the structure for slash-dot MusicXML element.
type SlashDot struct {
	XMLName xml.Name `xml:"slash-dot"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *SlashDot) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("slash-dot", strings.TrimSpace(r.IValue), attributes, children)
}
