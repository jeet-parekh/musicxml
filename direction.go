package musicxml

import (
	"encoding/xml"
	"strings"
)

// Direction is the structure for direction MusicXML element.
type Direction struct {
	XMLName xml.Name `xml:"direction"`

	DirectiveAttr string `xml:"directive,attr,omitempty"`
	IdAttr        string `xml:"id,attr,omitempty"`
	PlacementAttr string `xml:"placement,attr,omitempty"`

	DirectionType []DirectionType `xml:"direction-type,omitempty"`
	Footnote      []Footnote      `xml:"footnote,omitempty"`
	Level         []Level         `xml:"level,omitempty"`
	Offset        []Offset        `xml:"offset,omitempty"`
	Sound         []Sound         `xml:"sound,omitempty"`
	Staff         []Staff         `xml:"staff,omitempty"`
	Voice         []Voice         `xml:"voice,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Direction) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["directive"] = r.DirectiveAttr
	attributes["id"] = r.IdAttr
	attributes["placement"] = r.PlacementAttr

	children["direction-type"] = make([]*MXML, len(r.DirectionType))
	for i, c := range r.DirectionType {
		children["direction-type"][i] = c.ToMXML()
	}

	children["footnote"] = make([]*MXML, len(r.Footnote))
	for i, c := range r.Footnote {
		children["footnote"][i] = c.ToMXML()
	}

	children["level"] = make([]*MXML, len(r.Level))
	for i, c := range r.Level {
		children["level"][i] = c.ToMXML()
	}

	children["offset"] = make([]*MXML, len(r.Offset))
	for i, c := range r.Offset {
		children["offset"][i] = c.ToMXML()
	}

	children["sound"] = make([]*MXML, len(r.Sound))
	for i, c := range r.Sound {
		children["sound"][i] = c.ToMXML()
	}

	children["staff"] = make([]*MXML, len(r.Staff))
	for i, c := range r.Staff {
		children["staff"][i] = c.ToMXML()
	}

	children["voice"] = make([]*MXML, len(r.Voice))
	for i, c := range r.Voice {
		children["voice"][i] = c.ToMXML()
	}

	return newMXML("direction", strings.TrimSpace(r.IValue), attributes, children)
}
