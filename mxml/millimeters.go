package mxml

import (
	"encoding/xml"
	"strings"
)

// Millimeters is the structure for millimeters MusicXML element.
type Millimeters struct {
	XMLName xml.Name `xml:"millimeters"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Millimeters) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("millimeters", strings.TrimSpace(r.IValue), attributes, children)
}
