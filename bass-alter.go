package musicxml

import (
	"encoding/xml"
	"strings"
)

// BassAlter is the structure for bass-alter MusicXML element.
type BassAlter struct {
	XMLName xml.Name `xml:"bass-alter"`

	ColorAttr       string `xml:"color,attr,omitempty"`
	DefaultXAttr    string `xml:"default-x,attr,omitempty"`
	DefaultYAttr    string `xml:"default-y,attr,omitempty"`
	FontFamilyAttr  string `xml:"font-family,attr,omitempty"`
	FontSizeAttr    string `xml:"font-size,attr,omitempty"`
	FontStyleAttr   string `xml:"font-style,attr,omitempty"`
	FontWeightAttr  string `xml:"font-weight,attr,omitempty"`
	LocationAttr    string `xml:"location,attr,omitempty"`
	PrintObjectAttr string `xml:"print-object,attr,omitempty"`
	RelativeXAttr   string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr   string `xml:"relative-y,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *BassAlter) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["location"] = r.LocationAttr
	attributes["print-object"] = r.PrintObjectAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr

	return newMXML("bass-alter", strings.TrimSpace(r.IValue), attributes, children)
}
