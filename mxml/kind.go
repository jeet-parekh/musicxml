package mxml

import (
	"encoding/xml"
	"strings"
)

// Kind is the structure for kind MusicXML element.
type Kind struct {
	XMLName xml.Name `xml:"kind"`

	BracketDegreesAttr     string `xml:"bracket-degrees,attr,omitempty"`
	ColorAttr              string `xml:"color,attr,omitempty"`
	DefaultXAttr           string `xml:"default-x,attr,omitempty"`
	DefaultYAttr           string `xml:"default-y,attr,omitempty"`
	FontFamilyAttr         string `xml:"font-family,attr,omitempty"`
	FontSizeAttr           string `xml:"font-size,attr,omitempty"`
	FontStyleAttr          string `xml:"font-style,attr,omitempty"`
	FontWeightAttr         string `xml:"font-weight,attr,omitempty"`
	HalignAttr             string `xml:"halign,attr,omitempty"`
	ParenthesesDegreesAttr string `xml:"parentheses-degrees,attr,omitempty"`
	RelativeXAttr          string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr          string `xml:"relative-y,attr,omitempty"`
	StackDegreesAttr       string `xml:"stack-degrees,attr,omitempty"`
	TextAttr               string `xml:"text,attr,omitempty"`
	UseSymbolsAttr         string `xml:"use-symbols,attr,omitempty"`
	ValignAttr             string `xml:"valign,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Kind) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["bracket-degrees"] = r.BracketDegreesAttr
	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["halign"] = r.HalignAttr
	attributes["parentheses-degrees"] = r.ParenthesesDegreesAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["stack-degrees"] = r.StackDegreesAttr
	attributes["text"] = r.TextAttr
	attributes["use-symbols"] = r.UseSymbolsAttr
	attributes["valign"] = r.ValignAttr

	return newMXML("kind", strings.TrimSpace(r.IValue), attributes, children)
}
