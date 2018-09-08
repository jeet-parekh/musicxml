package mxml

import (
	"encoding/xml"
	"strings"
)

// Capo is the structure for capo MusicXML element.
type Capo struct {
	XMLName xml.Name `xml:"capo"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Capo) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("capo", strings.TrimSpace(r.IValue), attributes, children)
}
