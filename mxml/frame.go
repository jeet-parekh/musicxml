package mxml

import (
	"encoding/xml"
	"strings"
)

// Frame is the structure for frame MusicXML element.
type Frame struct {
	XMLName xml.Name `xml:"frame"`

	ColorAttr     string `xml:"color,attr,omitempty"`
	DefaultXAttr  string `xml:"default-x,attr,omitempty"`
	DefaultYAttr  string `xml:"default-y,attr,omitempty"`
	HalignAttr    string `xml:"halign,attr,omitempty"`
	HeightAttr    string `xml:"height,attr,omitempty"`
	IdAttr        string `xml:"id,attr,omitempty"`
	RelativeXAttr string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr string `xml:"relative-y,attr,omitempty"`
	UnplayedAttr  string `xml:"unplayed,attr,omitempty"`
	ValignAttr    string `xml:"valign,attr,omitempty"`
	WidthAttr     string `xml:"width,attr,omitempty"`

	FirstFret    []FirstFret    `xml:"first-fret,omitempty"`
	FrameFrets   []FrameFrets   `xml:"frame-frets,omitempty"`
	FrameNote    []FrameNote    `xml:"frame-note,omitempty"`
	FrameStrings []FrameStrings `xml:"frame-strings,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Frame) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["halign"] = r.HalignAttr
	attributes["height"] = r.HeightAttr
	attributes["id"] = r.IdAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["unplayed"] = r.UnplayedAttr
	attributes["valign"] = r.ValignAttr
	attributes["width"] = r.WidthAttr

	children["first-fret"] = make([]*MXML, len(r.FirstFret))
	for i, c := range r.FirstFret {
		children["first-fret"][i] = c.ToMXML()
	}

	children["frame-frets"] = make([]*MXML, len(r.FrameFrets))
	for i, c := range r.FrameFrets {
		children["frame-frets"][i] = c.ToMXML()
	}

	children["frame-note"] = make([]*MXML, len(r.FrameNote))
	for i, c := range r.FrameNote {
		children["frame-note"][i] = c.ToMXML()
	}

	children["frame-strings"] = make([]*MXML, len(r.FrameStrings))
	for i, c := range r.FrameStrings {
		children["frame-strings"][i] = c.ToMXML()
	}

	return newMXML("frame", strings.TrimSpace(r.IValue), attributes, children)
}
