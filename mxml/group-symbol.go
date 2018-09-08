package mxml

import (
	"encoding/xml"
	"strings"
)

// GroupSymbol is the structure for group-symbol MusicXML element.
type GroupSymbol struct {
	XMLName xml.Name `xml:"group-symbol"`

	ColorAttr     string `xml:"color,attr,omitempty"`
	DefaultXAttr  string `xml:"default-x,attr,omitempty"`
	DefaultYAttr  string `xml:"default-y,attr,omitempty"`
	RelativeXAttr string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr string `xml:"relative-y,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *GroupSymbol) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr

	return newMXML("group-symbol", strings.TrimSpace(r.IValue), attributes, children)
}
