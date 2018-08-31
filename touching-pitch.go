package musicxml

import (
	"encoding/xml"
	"strings"
)

// TouchingPitch is the structure for touching-pitch MusicXML element.
type TouchingPitch struct {
	XMLName xml.Name `xml:"touching-pitch"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *TouchingPitch) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("touching-pitch", strings.TrimSpace(r.IValue), attributes, children)
}
