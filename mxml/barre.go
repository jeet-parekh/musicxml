package mxml

import (
	"encoding/xml"
	"strings"
)

// Barre is the structure for barre MusicXML element.
type Barre struct {
	XMLName xml.Name `xml:"barre"`

	ColorAttr string `xml:"color,attr,omitempty"`
	TypeAttr  string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Barre) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["type"] = r.TypeAttr

	return newMXML("barre", strings.TrimSpace(r.IValue), attributes, children)
}
