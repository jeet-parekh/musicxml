package musicxml

import (
	"encoding/xml"
	"strings"
)

// Staccatissimo is the structure for staccatissimo MusicXML element.
type Staccatissimo struct {
	XMLName xml.Name `xml:"staccatissimo"`

	ColorAttr      string `xml:"color,attr,omitempty"`
	DefaultXAttr   string `xml:"default-x,attr,omitempty"`
	DefaultYAttr   string `xml:"default-y,attr,omitempty"`
	FontFamilyAttr string `xml:"font-family,attr,omitempty"`
	FontSizeAttr   string `xml:"font-size,attr,omitempty"`
	FontStyleAttr  string `xml:"font-style,attr,omitempty"`
	FontWeightAttr string `xml:"font-weight,attr,omitempty"`
	PlacementAttr  string `xml:"placement,attr,omitempty"`
	RelativeXAttr  string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr  string `xml:"relative-y,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Staccatissimo) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["placement"] = r.PlacementAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr

	return newMXML("staccatissimo", strings.TrimSpace(r.IValue), attributes, children)
}
