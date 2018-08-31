package musicxml

import (
	"encoding/xml"
	"strings"
)

// Divisions is the structure for divisions MusicXML element.
type Divisions struct {
	XMLName xml.Name `xml:"divisions"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Divisions) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("divisions", strings.TrimSpace(r.IValue), attributes, children)
}
