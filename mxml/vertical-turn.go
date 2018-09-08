package mxml

import (
	"encoding/xml"
	"strings"
)

// VerticalTurn is the structure for vertical-turn MusicXML element.
type VerticalTurn struct {
	XMLName xml.Name `xml:"vertical-turn"`

	AccelerateAttr  string `xml:"accelerate,attr,omitempty"`
	BeatsAttr       string `xml:"beats,attr,omitempty"`
	ColorAttr       string `xml:"color,attr,omitempty"`
	DefaultXAttr    string `xml:"default-x,attr,omitempty"`
	DefaultYAttr    string `xml:"default-y,attr,omitempty"`
	FontFamilyAttr  string `xml:"font-family,attr,omitempty"`
	FontSizeAttr    string `xml:"font-size,attr,omitempty"`
	FontStyleAttr   string `xml:"font-style,attr,omitempty"`
	FontWeightAttr  string `xml:"font-weight,attr,omitempty"`
	LastBeatAttr    string `xml:"last-beat,attr,omitempty"`
	PlacementAttr   string `xml:"placement,attr,omitempty"`
	RelativeXAttr   string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr   string `xml:"relative-y,attr,omitempty"`
	SecondBeatAttr  string `xml:"second-beat,attr,omitempty"`
	StartNoteAttr   string `xml:"start-note,attr,omitempty"`
	TrillStepAttr   string `xml:"trill-step,attr,omitempty"`
	TwoNoteTurnAttr string `xml:"two-note-turn,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *VerticalTurn) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["accelerate"] = r.AccelerateAttr
	attributes["beats"] = r.BeatsAttr
	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["last-beat"] = r.LastBeatAttr
	attributes["placement"] = r.PlacementAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["second-beat"] = r.SecondBeatAttr
	attributes["start-note"] = r.StartNoteAttr
	attributes["trill-step"] = r.TrillStepAttr
	attributes["two-note-turn"] = r.TwoNoteTurnAttr

	return newMXML("vertical-turn", strings.TrimSpace(r.IValue), attributes, children)
}
