package musicxml

import (
	"encoding/xml"
	"strings"
)

// KeyStep is the structure for key-step MusicXML element.
type KeyStep struct {
	XMLName xml.Name `xml:"key-step"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *KeyStep) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("key-step", strings.TrimSpace(r.IValue), attributes, children)
}
