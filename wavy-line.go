package musicxml

import (
	"encoding/xml"
	"strings"
)

// WavyLine is the structure for wavy-line MusicXML element.
type WavyLine struct {
	XMLName xml.Name `xml:"wavy-line"`

	AccelerateAttr  string `xml:"accelerate,attr,omitempty"`
	BeatsAttr       string `xml:"beats,attr,omitempty"`
	ColorAttr       string `xml:"color,attr,omitempty"`
	DefaultXAttr    string `xml:"default-x,attr,omitempty"`
	DefaultYAttr    string `xml:"default-y,attr,omitempty"`
	LastBeatAttr    string `xml:"last-beat,attr,omitempty"`
	NumberAttr      string `xml:"number,attr,omitempty"`
	PlacementAttr   string `xml:"placement,attr,omitempty"`
	RelativeXAttr   string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr   string `xml:"relative-y,attr,omitempty"`
	SecondBeatAttr  string `xml:"second-beat,attr,omitempty"`
	StartNoteAttr   string `xml:"start-note,attr,omitempty"`
	TrillStepAttr   string `xml:"trill-step,attr,omitempty"`
	TwoNoteTurnAttr string `xml:"two-note-turn,attr,omitempty"`
	TypeAttr        string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *WavyLine) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["accelerate"] = r.AccelerateAttr
	attributes["beats"] = r.BeatsAttr
	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["last-beat"] = r.LastBeatAttr
	attributes["number"] = r.NumberAttr
	attributes["placement"] = r.PlacementAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["second-beat"] = r.SecondBeatAttr
	attributes["start-note"] = r.StartNoteAttr
	attributes["trill-step"] = r.TrillStepAttr
	attributes["two-note-turn"] = r.TwoNoteTurnAttr
	attributes["type"] = r.TypeAttr

	return newMXML("wavy-line", strings.TrimSpace(r.IValue), attributes, children)
}
