package mxml

import (
	"encoding/xml"
	"strings"
)

// Harmony is the structure for harmony MusicXML element.
type Harmony struct {
	XMLName xml.Name `xml:"harmony"`

	ColorAttr       string `xml:"color,attr,omitempty"`
	DefaultXAttr    string `xml:"default-x,attr,omitempty"`
	DefaultYAttr    string `xml:"default-y,attr,omitempty"`
	FontFamilyAttr  string `xml:"font-family,attr,omitempty"`
	FontSizeAttr    string `xml:"font-size,attr,omitempty"`
	FontStyleAttr   string `xml:"font-style,attr,omitempty"`
	FontWeightAttr  string `xml:"font-weight,attr,omitempty"`
	IdAttr          string `xml:"id,attr,omitempty"`
	PlacementAttr   string `xml:"placement,attr,omitempty"`
	PrintFrameAttr  string `xml:"print-frame,attr,omitempty"`
	PrintObjectAttr string `xml:"print-object,attr,omitempty"`
	RelativeXAttr   string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr   string `xml:"relative-y,attr,omitempty"`
	TypeAttr        string `xml:"type,attr,omitempty"`

	Bass      []Bass      `xml:"bass,omitempty"`
	Degree    []Degree    `xml:"degree,omitempty"`
	Footnote  []Footnote  `xml:"footnote,omitempty"`
	Frame     []Frame     `xml:"frame,omitempty"`
	Function  []Function  `xml:"function,omitempty"`
	Inversion []Inversion `xml:"inversion,omitempty"`
	Kind      []Kind      `xml:"kind,omitempty"`
	Level     []Level     `xml:"level,omitempty"`
	Offset    []Offset    `xml:"offset,omitempty"`
	Root      []Root      `xml:"root,omitempty"`
	Staff     []Staff     `xml:"staff,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Harmony) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["id"] = r.IdAttr
	attributes["placement"] = r.PlacementAttr
	attributes["print-frame"] = r.PrintFrameAttr
	attributes["print-object"] = r.PrintObjectAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["type"] = r.TypeAttr

	children["bass"] = make([]*MXML, len(r.Bass))
	for i, c := range r.Bass {
		children["bass"][i] = c.ToMXML()
	}

	children["degree"] = make([]*MXML, len(r.Degree))
	for i, c := range r.Degree {
		children["degree"][i] = c.ToMXML()
	}

	children["footnote"] = make([]*MXML, len(r.Footnote))
	for i, c := range r.Footnote {
		children["footnote"][i] = c.ToMXML()
	}

	children["frame"] = make([]*MXML, len(r.Frame))
	for i, c := range r.Frame {
		children["frame"][i] = c.ToMXML()
	}

	children["function"] = make([]*MXML, len(r.Function))
	for i, c := range r.Function {
		children["function"][i] = c.ToMXML()
	}

	children["inversion"] = make([]*MXML, len(r.Inversion))
	for i, c := range r.Inversion {
		children["inversion"][i] = c.ToMXML()
	}

	children["kind"] = make([]*MXML, len(r.Kind))
	for i, c := range r.Kind {
		children["kind"][i] = c.ToMXML()
	}

	children["level"] = make([]*MXML, len(r.Level))
	for i, c := range r.Level {
		children["level"][i] = c.ToMXML()
	}

	children["offset"] = make([]*MXML, len(r.Offset))
	for i, c := range r.Offset {
		children["offset"][i] = c.ToMXML()
	}

	children["root"] = make([]*MXML, len(r.Root))
	for i, c := range r.Root {
		children["root"][i] = c.ToMXML()
	}

	children["staff"] = make([]*MXML, len(r.Staff))
	for i, c := range r.Staff {
		children["staff"][i] = c.ToMXML()
	}

	return newMXML("harmony", strings.TrimSpace(r.IValue), attributes, children)
}
