package musicxml

import (
	"encoding/xml"
	"strings"
)

// ExceptVoice is the structure for except-voice MusicXML element.
type ExceptVoice struct {
	XMLName xml.Name `xml:"except-voice"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *ExceptVoice) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("except-voice", strings.TrimSpace(r.IValue), attributes, children)
}
