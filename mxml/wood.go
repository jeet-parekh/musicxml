package mxml

import (
	"encoding/xml"
	"strings"
)

// Wood is the structure for wood MusicXML element.
type Wood struct {
	XMLName xml.Name `xml:"wood"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Wood) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("wood", strings.TrimSpace(r.IValue), attributes, children)
}
