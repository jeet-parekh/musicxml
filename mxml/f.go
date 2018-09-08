package mxml

import (
	"encoding/xml"
	"strings"
)

// F is the structure for f MusicXML element.
type F struct {
	XMLName xml.Name `xml:"f"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *F) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("f", strings.TrimSpace(r.IValue), attributes, children)
}
