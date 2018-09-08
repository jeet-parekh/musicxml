package mxml

import (
	"encoding/xml"
	"strings"
)

// BeatType is the structure for beat-type MusicXML element.
type BeatType struct {
	XMLName xml.Name `xml:"beat-type"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *BeatType) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("beat-type", strings.TrimSpace(r.IValue), attributes, children)
}
