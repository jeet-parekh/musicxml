package musicxml

import (
	"encoding/xml"
	"strings"
)

// KeyOctave is the structure for key-octave MusicXML element.
type KeyOctave struct {
	XMLName xml.Name `xml:"key-octave"`

	CancelAttr string `xml:"cancel,attr,omitempty"`
	NumberAttr string `xml:"number,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *KeyOctave) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["cancel"] = r.CancelAttr
	attributes["number"] = r.NumberAttr

	return newMXML("key-octave", strings.TrimSpace(r.IValue), attributes, children)
}
