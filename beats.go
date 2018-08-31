package musicxml

import (
	"encoding/xml"
	"strings"
)

// Beats is the structure for beats MusicXML element.
type Beats struct {
	XMLName xml.Name `xml:"beats"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Beats) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("beats", strings.TrimSpace(r.IValue), attributes, children)
}
