package musicxml

import (
	"encoding/xml"
	"strings"
)

// ClefOctaveChange is the structure for clef-octave-change MusicXML element.
type ClefOctaveChange struct {
	XMLName xml.Name `xml:"clef-octave-change"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *ClefOctaveChange) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("clef-octave-change", strings.TrimSpace(r.IValue), attributes, children)
}
