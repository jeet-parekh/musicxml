package musicxml

import (
	"encoding/xml"
	"strings"
)

// SystemDistance is the structure for system-distance MusicXML element.
type SystemDistance struct {
	XMLName xml.Name `xml:"system-distance"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *SystemDistance) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("system-distance", strings.TrimSpace(r.IValue), attributes, children)
}
