package mxml

import (
	"encoding/xml"
	"strings"
)

// Sfp is the structure for sfp MusicXML element.
type Sfp struct {
	XMLName xml.Name `xml:"sfp"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Sfp) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("sfp", strings.TrimSpace(r.IValue), attributes, children)
}
