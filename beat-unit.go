package musicxml

import (
	"encoding/xml"
	"strings"
)

// BeatUnit is the structure for beat-unit MusicXML element.
type BeatUnit struct {
	XMLName xml.Name `xml:"beat-unit"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *BeatUnit) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("beat-unit", strings.TrimSpace(r.IValue), attributes, children)
}
