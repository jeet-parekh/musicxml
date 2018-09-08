package mxml

import (
	"encoding/xml"
	"strings"
)

// PartSymbol is the structure for part-symbol MusicXML element.
type PartSymbol struct {
	XMLName xml.Name `xml:"part-symbol"`

	BottomStaffAttr string `xml:"bottom-staff,attr,omitempty"`
	ColorAttr       string `xml:"color,attr,omitempty"`
	DefaultXAttr    string `xml:"default-x,attr,omitempty"`
	DefaultYAttr    string `xml:"default-y,attr,omitempty"`
	RelativeXAttr   string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr   string `xml:"relative-y,attr,omitempty"`
	TopStaffAttr    string `xml:"top-staff,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *PartSymbol) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["bottom-staff"] = r.BottomStaffAttr
	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["top-staff"] = r.TopStaffAttr

	return newMXML("part-symbol", strings.TrimSpace(r.IValue), attributes, children)
}
