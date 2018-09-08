package mxml

import (
	"encoding/xml"
	"strings"
)

// TupletNumber is the structure for tuplet-number MusicXML element.
type TupletNumber struct {
	XMLName xml.Name `xml:"tuplet-number"`

	ColorAttr      string `xml:"color,attr,omitempty"`
	FontFamilyAttr string `xml:"font-family,attr,omitempty"`
	FontSizeAttr   string `xml:"font-size,attr,omitempty"`
	FontStyleAttr  string `xml:"font-style,attr,omitempty"`
	FontWeightAttr string `xml:"font-weight,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *TupletNumber) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr

	return newMXML("tuplet-number", strings.TrimSpace(r.IValue), attributes, children)
}
