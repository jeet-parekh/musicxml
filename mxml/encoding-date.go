package mxml

import (
	"encoding/xml"
	"strings"
)

// EncodingDate is the structure for encoding-date MusicXML element.
type EncodingDate struct {
	XMLName xml.Name `xml:"encoding-date"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *EncodingDate) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("encoding-date", strings.TrimSpace(r.IValue), attributes, children)
}
