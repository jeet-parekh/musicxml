package mxml

import (
	"encoding/xml"
	"strings"
)

// BeatUnitDot is the structure for beat-unit-dot MusicXML element.
type BeatUnitDot struct {
	XMLName xml.Name `xml:"beat-unit-dot"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *BeatUnitDot) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("beat-unit-dot", strings.TrimSpace(r.IValue), attributes, children)
}
