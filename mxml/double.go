package mxml

import (
	"encoding/xml"
	"strings"
)

// Double is the structure for double MusicXML element.
type Double struct {
	XMLName xml.Name `xml:"double"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Double) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("double", strings.TrimSpace(r.IValue), attributes, children)
}
