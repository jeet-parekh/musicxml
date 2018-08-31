package musicxml

import (
	"encoding/xml"
	"strings"
)

// Arpeggiate is the structure for arpeggiate MusicXML element.
type Arpeggiate struct {
	XMLName xml.Name `xml:"arpeggiate"`

	ColorAttr     string `xml:"color,attr,omitempty"`
	DefaultXAttr  string `xml:"default-x,attr,omitempty"`
	DefaultYAttr  string `xml:"default-y,attr,omitempty"`
	DirectionAttr string `xml:"direction,attr,omitempty"`
	IdAttr        string `xml:"id,attr,omitempty"`
	NumberAttr    string `xml:"number,attr,omitempty"`
	PlacementAttr string `xml:"placement,attr,omitempty"`
	RelativeXAttr string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr string `xml:"relative-y,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Arpeggiate) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["direction"] = r.DirectionAttr
	attributes["id"] = r.IdAttr
	attributes["number"] = r.NumberAttr
	attributes["placement"] = r.PlacementAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr

	return newMXML("arpeggiate", strings.TrimSpace(r.IValue), attributes, children)
}
