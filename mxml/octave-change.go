package mxml

import (
	"encoding/xml"
	"strings"
)

// OctaveChange is the structure for octave-change MusicXML element.
type OctaveChange struct {
	XMLName xml.Name `xml:"octave-change"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *OctaveChange) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("octave-change", strings.TrimSpace(r.IValue), attributes, children)
}
