package mxml

import (
	"encoding/xml"
	"strings"
)

// HarmonMute is the structure for harmon-mute MusicXML element.
type HarmonMute struct {
	XMLName xml.Name `xml:"harmon-mute"`

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

	HarmonClosed []HarmonClosed `xml:"harmon-closed,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *HarmonMute) ToMXML() *MXML {
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

	children["harmon-closed"] = make([]*MXML, len(r.HarmonClosed))
	for i, c := range r.HarmonClosed {
		children["harmon-closed"][i] = c.ToMXML()
	}

	return newMXML("harmon-mute", strings.TrimSpace(r.IValue), attributes, children)
}
