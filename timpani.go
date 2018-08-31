package musicxml

import (
	"encoding/xml"
	"strings"
)

// Timpani is the structure for timpani MusicXML element.
type Timpani struct {
	XMLName xml.Name `xml:"timpani"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Timpani) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("timpani", strings.TrimSpace(r.IValue), attributes, children)
}
