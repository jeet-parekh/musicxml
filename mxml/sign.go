package mxml

import (
	"encoding/xml"
	"strings"
)

// Sign is the structure for sign MusicXML element.
type Sign struct {
	XMLName xml.Name `xml:"sign"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Sign) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("sign", strings.TrimSpace(r.IValue), attributes, children)
}
