package musicxml

import (
	"encoding/xml"
	"strings"
)

// DisplayOctave is the structure for display-octave MusicXML element.
type DisplayOctave struct {
	XMLName xml.Name `xml:"display-octave"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *DisplayOctave) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("display-octave", strings.TrimSpace(r.IValue), attributes, children)
}
