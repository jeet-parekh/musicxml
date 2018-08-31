package musicxml

import (
	"encoding/xml"
	"strings"
)

// Group is the structure for group MusicXML element.
type Group struct {
	XMLName xml.Name `xml:"group"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Group) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("group", strings.TrimSpace(r.IValue), attributes, children)
}
