package musicxml

import (
	"encoding/xml"
	"strings"
)

// Elision is the structure for elision MusicXML element.
type Elision struct {
	XMLName xml.Name `xml:"elision"`

	ColorAttr      string `xml:"color,attr,omitempty"`
	FontFamilyAttr string `xml:"font-family,attr,omitempty"`
	FontSizeAttr   string `xml:"font-size,attr,omitempty"`
	FontStyleAttr  string `xml:"font-style,attr,omitempty"`
	FontWeightAttr string `xml:"font-weight,attr,omitempty"`
	SmuflAttr      string `xml:"smufl,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Elision) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["smufl"] = r.SmuflAttr

	return newMXML("elision", strings.TrimSpace(r.IValue), attributes, children)
}
