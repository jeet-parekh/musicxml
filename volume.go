package musicxml

import (
	"encoding/xml"
	"strings"
)

// Volume is the structure for volume MusicXML element.
type Volume struct {
	XMLName xml.Name `xml:"volume"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Volume) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("volume", strings.TrimSpace(r.IValue), attributes, children)
}
