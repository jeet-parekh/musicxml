package musicxml

import (
	"encoding/xml"
	"strings"
)

// Fp is the structure for fp MusicXML element.
type Fp struct {
	XMLName xml.Name `xml:"fp"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Fp) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("fp", strings.TrimSpace(r.IValue), attributes, children)
}
