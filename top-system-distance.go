package musicxml

import (
	"encoding/xml"
	"strings"
)

// TopSystemDistance is the structure for top-system-distance MusicXML element.
type TopSystemDistance struct {
	XMLName xml.Name `xml:"top-system-distance"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *TopSystemDistance) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("top-system-distance", strings.TrimSpace(r.IValue), attributes, children)
}
