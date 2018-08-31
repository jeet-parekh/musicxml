package musicxml

import (
	"encoding/xml"
	"strings"
)

// HoleShape is the structure for hole-shape MusicXML element.
type HoleShape struct {
	XMLName xml.Name `xml:"hole-shape"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *HoleShape) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("hole-shape", strings.TrimSpace(r.IValue), attributes, children)
}
