package musicxml

import (
	"encoding/xml"
	"strings"
)

// Sfzp is the structure for sfzp MusicXML element.
type Sfzp struct {
	XMLName xml.Name `xml:"sfzp"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Sfzp) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("sfzp", strings.TrimSpace(r.IValue), attributes, children)
}
