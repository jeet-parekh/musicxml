package mxml

import (
	"encoding/xml"
	"strings"
)

// Elevation is the structure for elevation MusicXML element.
type Elevation struct {
	XMLName xml.Name `xml:"elevation"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Elevation) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("elevation", strings.TrimSpace(r.IValue), attributes, children)
}
