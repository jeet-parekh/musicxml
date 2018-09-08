package mxml

import (
	"encoding/xml"
	"strings"
)

// Scoop is the structure for scoop MusicXML element.
type Scoop struct {
	XMLName xml.Name `xml:"scoop"`

	ColorAttr       string `xml:"color,attr,omitempty"`
	DashLengthAttr  string `xml:"dash-length,attr,omitempty"`
	DefaultXAttr    string `xml:"default-x,attr,omitempty"`
	DefaultYAttr    string `xml:"default-y,attr,omitempty"`
	FontFamilyAttr  string `xml:"font-family,attr,omitempty"`
	FontSizeAttr    string `xml:"font-size,attr,omitempty"`
	FontStyleAttr   string `xml:"font-style,attr,omitempty"`
	FontWeightAttr  string `xml:"font-weight,attr,omitempty"`
	LineLengthAttr  string `xml:"line-length,attr,omitempty"`
	LineShapeAttr   string `xml:"line-shape,attr,omitempty"`
	LineTypeAttr    string `xml:"line-type,attr,omitempty"`
	PlacementAttr   string `xml:"placement,attr,omitempty"`
	RelativeXAttr   string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr   string `xml:"relative-y,attr,omitempty"`
	SpaceLengthAttr string `xml:"space-length,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Scoop) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["dash-length"] = r.DashLengthAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["line-length"] = r.LineLengthAttr
	attributes["line-shape"] = r.LineShapeAttr
	attributes["line-type"] = r.LineTypeAttr
	attributes["placement"] = r.PlacementAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["space-length"] = r.SpaceLengthAttr

	return newMXML("scoop", strings.TrimSpace(r.IValue), attributes, children)
}
