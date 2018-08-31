package musicxml

import (
	"encoding/xml"
	"strings"
)

// PerMinute is the structure for per-minute MusicXML element.
type PerMinute struct {
	XMLName xml.Name `xml:"per-minute"`

	FontFamilyAttr string `xml:"font-family,attr,omitempty"`
	FontSizeAttr   string `xml:"font-size,attr,omitempty"`
	FontStyleAttr  string `xml:"font-style,attr,omitempty"`
	FontWeightAttr string `xml:"font-weight,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *PerMinute) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr

	return newMXML("per-minute", strings.TrimSpace(r.IValue), attributes, children)
}
