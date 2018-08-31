package musicxml

import (
	"encoding/xml"
	"strings"
)

// Hole is the structure for hole MusicXML element.
type Hole struct {
	XMLName xml.Name `xml:"hole"`

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

	HoleClosed []HoleClosed `xml:"hole-closed,omitempty"`
	HoleShape  []HoleShape  `xml:"hole-shape,omitempty"`
	HoleType   []HoleType   `xml:"hole-type,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Hole) ToMXML() *MXML {
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

	children["hole-closed"] = make([]*MXML, len(r.HoleClosed))
	for i, c := range r.HoleClosed {
		children["hole-closed"][i] = c.ToMXML()
	}

	children["hole-shape"] = make([]*MXML, len(r.HoleShape))
	for i, c := range r.HoleShape {
		children["hole-shape"][i] = c.ToMXML()
	}

	children["hole-type"] = make([]*MXML, len(r.HoleType))
	for i, c := range r.HoleType {
		children["hole-type"][i] = c.ToMXML()
	}

	return newMXML("hole", strings.TrimSpace(r.IValue), attributes, children)
}
