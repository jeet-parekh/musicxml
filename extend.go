package musicxml

import (
	"encoding/xml"
	"strings"
)

// Extend is the structure for extend MusicXML element.
type Extend struct {
	XMLName xml.Name `xml:"extend"`

	ColorAttr     string `xml:"color,attr,omitempty"`
	DefaultXAttr  string `xml:"default-x,attr,omitempty"`
	DefaultYAttr  string `xml:"default-y,attr,omitempty"`
	RelativeXAttr string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr string `xml:"relative-y,attr,omitempty"`
	TypeAttr      string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Extend) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["type"] = r.TypeAttr

	return newMXML("extend", strings.TrimSpace(r.IValue), attributes, children)
}
