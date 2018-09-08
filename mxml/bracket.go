package mxml

import (
	"encoding/xml"
	"strings"
)

// Bracket is the structure for bracket MusicXML element.
type Bracket struct {
	XMLName xml.Name `xml:"bracket"`

	ColorAttr       string `xml:"color,attr,omitempty"`
	DashLengthAttr  string `xml:"dash-length,attr,omitempty"`
	DefaultXAttr    string `xml:"default-x,attr,omitempty"`
	DefaultYAttr    string `xml:"default-y,attr,omitempty"`
	EndLengthAttr   string `xml:"end-length,attr,omitempty"`
	IdAttr          string `xml:"id,attr,omitempty"`
	LineEndAttr     string `xml:"line-end,attr,omitempty"`
	LineTypeAttr    string `xml:"line-type,attr,omitempty"`
	NumberAttr      string `xml:"number,attr,omitempty"`
	RelativeXAttr   string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr   string `xml:"relative-y,attr,omitempty"`
	SpaceLengthAttr string `xml:"space-length,attr,omitempty"`
	TypeAttr        string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Bracket) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["dash-length"] = r.DashLengthAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["end-length"] = r.EndLengthAttr
	attributes["id"] = r.IdAttr
	attributes["line-end"] = r.LineEndAttr
	attributes["line-type"] = r.LineTypeAttr
	attributes["number"] = r.NumberAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["space-length"] = r.SpaceLengthAttr
	attributes["type"] = r.TypeAttr

	return newMXML("bracket", strings.TrimSpace(r.IValue), attributes, children)
}
