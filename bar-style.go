package musicxml

import (
	"encoding/xml"
	"strings"
)

// BarStyle is the structure for bar-style MusicXML element.
type BarStyle struct {
	XMLName xml.Name `xml:"bar-style"`

	ColorAttr string `xml:"color,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *BarStyle) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr

	return newMXML("bar-style", strings.TrimSpace(r.IValue), attributes, children)
}
