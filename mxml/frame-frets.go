package mxml

import (
	"encoding/xml"
	"strings"
)

// FrameFrets is the structure for frame-frets MusicXML element.
type FrameFrets struct {
	XMLName xml.Name `xml:"frame-frets"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *FrameFrets) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("frame-frets", strings.TrimSpace(r.IValue), attributes, children)
}
