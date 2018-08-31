package musicxml

import (
	"encoding/xml"
	"strings"
)

// KeyAlter is the structure for key-alter MusicXML element.
type KeyAlter struct {
	XMLName xml.Name `xml:"key-alter"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *KeyAlter) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("key-alter", strings.TrimSpace(r.IValue), attributes, children)
}
