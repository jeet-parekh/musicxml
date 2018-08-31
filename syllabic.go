package musicxml

import (
	"encoding/xml"
	"strings"
)

// Syllabic is the structure for syllabic MusicXML element.
type Syllabic struct {
	XMLName xml.Name `xml:"syllabic"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Syllabic) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("syllabic", strings.TrimSpace(r.IValue), attributes, children)
}
