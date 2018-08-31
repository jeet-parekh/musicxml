package musicxml

import (
	"encoding/xml"
	"strings"
)

// Slide is the structure for slide MusicXML element.
type Slide struct {
	XMLName xml.Name `xml:"slide"`

	AccelerateAttr  string `xml:"accelerate,attr,omitempty"`
	BeatsAttr       string `xml:"beats,attr,omitempty"`
	ColorAttr       string `xml:"color,attr,omitempty"`
	DashLengthAttr  string `xml:"dash-length,attr,omitempty"`
	DefaultXAttr    string `xml:"default-x,attr,omitempty"`
	DefaultYAttr    string `xml:"default-y,attr,omitempty"`
	FirstBeatAttr   string `xml:"first-beat,attr,omitempty"`
	FontFamilyAttr  string `xml:"font-family,attr,omitempty"`
	FontSizeAttr    string `xml:"font-size,attr,omitempty"`
	FontStyleAttr   string `xml:"font-style,attr,omitempty"`
	FontWeightAttr  string `xml:"font-weight,attr,omitempty"`
	IdAttr          string `xml:"id,attr,omitempty"`
	LastBeatAttr    string `xml:"last-beat,attr,omitempty"`
	LineTypeAttr    string `xml:"line-type,attr,omitempty"`
	NumberAttr      string `xml:"number,attr,omitempty"`
	RelativeXAttr   string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr   string `xml:"relative-y,attr,omitempty"`
	SpaceLengthAttr string `xml:"space-length,attr,omitempty"`
	TypeAttr        string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Slide) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["accelerate"] = r.AccelerateAttr
	attributes["beats"] = r.BeatsAttr
	attributes["color"] = r.ColorAttr
	attributes["dash-length"] = r.DashLengthAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["first-beat"] = r.FirstBeatAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["id"] = r.IdAttr
	attributes["last-beat"] = r.LastBeatAttr
	attributes["line-type"] = r.LineTypeAttr
	attributes["number"] = r.NumberAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["space-length"] = r.SpaceLengthAttr
	attributes["type"] = r.TypeAttr

	return newMXML("slide", strings.TrimSpace(r.IValue), attributes, children)
}
