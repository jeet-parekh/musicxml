package mxml

import (
	"encoding/xml"
	"strings"
)

// Harmonic is the structure for harmonic MusicXML element.
type Harmonic struct {
	XMLName xml.Name `xml:"harmonic"`

	ColorAttr       string `xml:"color,attr,omitempty"`
	DefaultXAttr    string `xml:"default-x,attr,omitempty"`
	DefaultYAttr    string `xml:"default-y,attr,omitempty"`
	FontFamilyAttr  string `xml:"font-family,attr,omitempty"`
	FontSizeAttr    string `xml:"font-size,attr,omitempty"`
	FontStyleAttr   string `xml:"font-style,attr,omitempty"`
	FontWeightAttr  string `xml:"font-weight,attr,omitempty"`
	PlacementAttr   string `xml:"placement,attr,omitempty"`
	PrintObjectAttr string `xml:"print-object,attr,omitempty"`
	RelativeXAttr   string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr   string `xml:"relative-y,attr,omitempty"`

	Artificial    []Artificial    `xml:"artificial,omitempty"`
	BasePitch     []BasePitch     `xml:"base-pitch,omitempty"`
	Natural       []Natural       `xml:"natural,omitempty"`
	SoundingPitch []SoundingPitch `xml:"sounding-pitch,omitempty"`
	TouchingPitch []TouchingPitch `xml:"touching-pitch,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Harmonic) ToMXML() *MXML {
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
	attributes["print-object"] = r.PrintObjectAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr

	children["artificial"] = make([]*MXML, len(r.Artificial))
	for i, c := range r.Artificial {
		children["artificial"][i] = c.ToMXML()
	}

	children["base-pitch"] = make([]*MXML, len(r.BasePitch))
	for i, c := range r.BasePitch {
		children["base-pitch"][i] = c.ToMXML()
	}

	children["natural"] = make([]*MXML, len(r.Natural))
	for i, c := range r.Natural {
		children["natural"][i] = c.ToMXML()
	}

	children["sounding-pitch"] = make([]*MXML, len(r.SoundingPitch))
	for i, c := range r.SoundingPitch {
		children["sounding-pitch"][i] = c.ToMXML()
	}

	children["touching-pitch"] = make([]*MXML, len(r.TouchingPitch))
	for i, c := range r.TouchingPitch {
		children["touching-pitch"][i] = c.ToMXML()
	}

	return newMXML("harmonic", strings.TrimSpace(r.IValue), attributes, children)
}
