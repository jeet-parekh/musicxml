package musicxml

import (
	"encoding/xml"
	"strings"
)

// Bend is the structure for bend MusicXML element.
type Bend struct {
	XMLName xml.Name `xml:"bend"`

	AccelerateAttr string `xml:"accelerate,attr,omitempty"`
	BeatsAttr      string `xml:"beats,attr,omitempty"`
	ColorAttr      string `xml:"color,attr,omitempty"`
	DefaultXAttr   string `xml:"default-x,attr,omitempty"`
	DefaultYAttr   string `xml:"default-y,attr,omitempty"`
	FirstBeatAttr  string `xml:"first-beat,attr,omitempty"`
	FontFamilyAttr string `xml:"font-family,attr,omitempty"`
	FontSizeAttr   string `xml:"font-size,attr,omitempty"`
	FontStyleAttr  string `xml:"font-style,attr,omitempty"`
	FontWeightAttr string `xml:"font-weight,attr,omitempty"`
	LastBeatAttr   string `xml:"last-beat,attr,omitempty"`
	RelativeXAttr  string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr  string `xml:"relative-y,attr,omitempty"`

	BendAlter []BendAlter `xml:"bend-alter,omitempty"`
	PreBend   []PreBend   `xml:"pre-bend,omitempty"`
	Release   []Release   `xml:"release,omitempty"`
	WithBar   []WithBar   `xml:"with-bar,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Bend) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["accelerate"] = r.AccelerateAttr
	attributes["beats"] = r.BeatsAttr
	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["first-beat"] = r.FirstBeatAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["last-beat"] = r.LastBeatAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr

	children["bend-alter"] = make([]*MXML, len(r.BendAlter))
	for i, c := range r.BendAlter {
		children["bend-alter"][i] = c.ToMXML()
	}

	children["pre-bend"] = make([]*MXML, len(r.PreBend))
	for i, c := range r.PreBend {
		children["pre-bend"][i] = c.ToMXML()
	}

	children["release"] = make([]*MXML, len(r.Release))
	for i, c := range r.Release {
		children["release"][i] = c.ToMXML()
	}

	children["with-bar"] = make([]*MXML, len(r.WithBar))
	for i, c := range r.WithBar {
		children["with-bar"][i] = c.ToMXML()
	}

	return newMXML("bend", strings.TrimSpace(r.IValue), attributes, children)
}
