package mxml

import (
	"encoding/xml"
	"strings"
)

// SoundingPitch is the structure for sounding-pitch MusicXML element.
type SoundingPitch struct {
	XMLName xml.Name `xml:"sounding-pitch"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *SoundingPitch) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("sounding-pitch", strings.TrimSpace(r.IValue), attributes, children)
}
