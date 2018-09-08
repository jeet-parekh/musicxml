package mxml

import (
	"encoding/xml"
	"strings"
)

// Mute is the structure for mute MusicXML element.
type Mute struct {
	XMLName xml.Name `xml:"mute"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Mute) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("mute", strings.TrimSpace(r.IValue), attributes, children)
}
