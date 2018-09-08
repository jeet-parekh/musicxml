package mxml

import (
	"encoding/xml"
	"strings"
)

// Cue is the structure for cue MusicXML element.
type Cue struct {
	XMLName xml.Name `xml:"cue"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Cue) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("cue", strings.TrimSpace(r.IValue), attributes, children)
}
