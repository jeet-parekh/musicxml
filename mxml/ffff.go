package mxml

import (
	"encoding/xml"
	"strings"
)

// Ffff is the structure for ffff MusicXML element.
type Ffff struct {
	XMLName xml.Name `xml:"ffff"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Ffff) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("ffff", strings.TrimSpace(r.IValue), attributes, children)
}
