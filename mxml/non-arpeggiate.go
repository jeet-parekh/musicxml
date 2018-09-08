package mxml

import (
	"encoding/xml"
	"strings"
)

// NonArpeggiate is the structure for non-arpeggiate MusicXML element.
type NonArpeggiate struct {
	XMLName xml.Name `xml:"non-arpeggiate"`

	ColorAttr     string `xml:"color,attr,omitempty"`
	DefaultXAttr  string `xml:"default-x,attr,omitempty"`
	DefaultYAttr  string `xml:"default-y,attr,omitempty"`
	IdAttr        string `xml:"id,attr,omitempty"`
	NumberAttr    string `xml:"number,attr,omitempty"`
	PlacementAttr string `xml:"placement,attr,omitempty"`
	RelativeXAttr string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr string `xml:"relative-y,attr,omitempty"`
	TypeAttr      string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *NonArpeggiate) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["id"] = r.IdAttr
	attributes["number"] = r.NumberAttr
	attributes["placement"] = r.PlacementAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["type"] = r.TypeAttr

	return newMXML("non-arpeggiate", strings.TrimSpace(r.IValue), attributes, children)
}
