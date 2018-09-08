package mxml

import (
	"encoding/xml"
	"strings"
)

// Arrow is the structure for arrow MusicXML element.
type Arrow struct {
	XMLName xml.Name `xml:"arrow"`

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
	SmuflAttr      string `xml:"smufl,attr,omitempty"`

	ArrowDirection []ArrowDirection `xml:"arrow-direction,omitempty"`
	ArrowStyle     []ArrowStyle     `xml:"arrow-style,omitempty"`
	Arrowhead      []Arrowhead      `xml:"arrowhead,omitempty"`
	CircularArrow  []CircularArrow  `xml:"circular-arrow,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Arrow) ToMXML() *MXML {
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
	attributes["smufl"] = r.SmuflAttr

	children["arrow-direction"] = make([]*MXML, len(r.ArrowDirection))
	for i, c := range r.ArrowDirection {
		children["arrow-direction"][i] = c.ToMXML()
	}

	children["arrow-style"] = make([]*MXML, len(r.ArrowStyle))
	for i, c := range r.ArrowStyle {
		children["arrow-style"][i] = c.ToMXML()
	}

	children["arrowhead"] = make([]*MXML, len(r.Arrowhead))
	for i, c := range r.Arrowhead {
		children["arrowhead"][i] = c.ToMXML()
	}

	children["circular-arrow"] = make([]*MXML, len(r.CircularArrow))
	for i, c := range r.CircularArrow {
		children["circular-arrow"][i] = c.ToMXML()
	}

	return newMXML("arrow", strings.TrimSpace(r.IValue), attributes, children)
}
