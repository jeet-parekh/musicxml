package mxml

import (
	"encoding/xml"
	"strings"
)

// Laughing is the structure for laughing MusicXML element.
type Laughing struct {
	XMLName xml.Name `xml:"laughing"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Laughing) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("laughing", strings.TrimSpace(r.IValue), attributes, children)
}
