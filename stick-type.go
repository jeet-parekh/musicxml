package musicxml

import (
	"encoding/xml"
	"strings"
)

// StickType is the structure for stick-type MusicXML element.
type StickType struct {
	XMLName xml.Name `xml:"stick-type"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *StickType) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("stick-type", strings.TrimSpace(r.IValue), attributes, children)
}
