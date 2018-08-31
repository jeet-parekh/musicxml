package musicxml

import (
	"encoding/xml"
	"strings"
)

// EncodingDescription is the structure for encoding-description MusicXML element.
type EncodingDescription struct {
	XMLName xml.Name `xml:"encoding-description"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *EncodingDescription) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("encoding-description", strings.TrimSpace(r.IValue), attributes, children)
}
