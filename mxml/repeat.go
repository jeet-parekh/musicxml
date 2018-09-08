package mxml

import (
	"encoding/xml"
	"strings"
)

// Repeat is the structure for repeat MusicXML element.
type Repeat struct {
	XMLName xml.Name `xml:"repeat"`

	DirectionAttr string `xml:"direction,attr,omitempty"`
	TimesAttr     string `xml:"times,attr,omitempty"`
	WingedAttr    string `xml:"winged,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Repeat) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["direction"] = r.DirectionAttr
	attributes["times"] = r.TimesAttr
	attributes["winged"] = r.WingedAttr

	return newMXML("repeat", strings.TrimSpace(r.IValue), attributes, children)
}
