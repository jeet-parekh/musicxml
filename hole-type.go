package musicxml

import (
	"encoding/xml"
	"strings"
)

// HoleType is the structure for hole-type MusicXML element.
type HoleType struct {
	XMLName xml.Name `xml:"hole-type"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *HoleType) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("hole-type", strings.TrimSpace(r.IValue), attributes, children)
}
