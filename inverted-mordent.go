package musicxml

import (
	"encoding/xml"
	"strings"
)

// InvertedMordent is the structure for inverted-mordent MusicXML element.
type InvertedMordent struct {
	XMLName xml.Name `xml:"inverted-mordent"`

	AccelerateAttr  string `xml:"accelerate,attr,omitempty"`
	ApproachAttr    string `xml:"approach,attr,omitempty"`
	BeatsAttr       string `xml:"beats,attr,omitempty"`
	ColorAttr       string `xml:"color,attr,omitempty"`
	DefaultXAttr    string `xml:"default-x,attr,omitempty"`
	DefaultYAttr    string `xml:"default-y,attr,omitempty"`
	DepartureAttr   string `xml:"departure,attr,omitempty"`
	FontFamilyAttr  string `xml:"font-family,attr,omitempty"`
	FontSizeAttr    string `xml:"font-size,attr,omitempty"`
	FontStyleAttr   string `xml:"font-style,attr,omitempty"`
	FontWeightAttr  string `xml:"font-weight,attr,omitempty"`
	LastBeatAttr    string `xml:"last-beat,attr,omitempty"`
	LongAttr        string `xml:"long,attr,omitempty"`
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
func (r *InvertedMordent) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["accelerate"] = r.AccelerateAttr
	attributes["approach"] = r.ApproachAttr
	attributes["beats"] = r.BeatsAttr
	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["departure"] = r.DepartureAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["last-beat"] = r.LastBeatAttr
	attributes["long"] = r.LongAttr
	attributes["placement"] = r.PlacementAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["second-beat"] = r.SecondBeatAttr
	attributes["start-note"] = r.StartNoteAttr
	attributes["trill-step"] = r.TrillStepAttr
	attributes["two-note-turn"] = r.TwoNoteTurnAttr

	return newMXML("inverted-mordent", strings.TrimSpace(r.IValue), attributes, children)
}
