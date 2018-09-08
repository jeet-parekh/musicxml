package mxml

import (
	"encoding/xml"
	"strings"
)

// Stick is the structure for stick MusicXML element.
type Stick struct {
	XMLName xml.Name `xml:"stick"`

	DashedCircleAttr string `xml:"dashed-circle,attr,omitempty"`
	ParenthesesAttr  string `xml:"parentheses,attr,omitempty"`
	TipAttr          string `xml:"tip,attr,omitempty"`

	StickMaterial []StickMaterial `xml:"stick-material,omitempty"`
	StickType     []StickType     `xml:"stick-type,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Stick) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["dashed-circle"] = r.DashedCircleAttr
	attributes["parentheses"] = r.ParenthesesAttr
	attributes["tip"] = r.TipAttr

	children["stick-material"] = make([]*MXML, len(r.StickMaterial))
	for i, c := range r.StickMaterial {
		children["stick-material"][i] = c.ToMXML()
	}

	children["stick-type"] = make([]*MXML, len(r.StickType))
	for i, c := range r.StickType {
		children["stick-type"][i] = c.ToMXML()
	}

	return newMXML("stick", strings.TrimSpace(r.IValue), attributes, children)
}
