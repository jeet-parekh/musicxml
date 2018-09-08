package mxml

import (
	"encoding/xml"
	"strings"
)

// Ornaments is the structure for ornaments MusicXML element.
type Ornaments struct {
	XMLName xml.Name `xml:"ornaments"`

	IdAttr string `xml:"id,attr,omitempty"`

	AccidentalMark       []AccidentalMark       `xml:"accidental-mark,omitempty"`
	DelayedInvertedTurn  []DelayedInvertedTurn  `xml:"delayed-inverted-turn,omitempty"`
	DelayedTurn          []DelayedTurn          `xml:"delayed-turn,omitempty"`
	Haydn                []Haydn                `xml:"haydn,omitempty"`
	InvertedMordent      []InvertedMordent      `xml:"inverted-mordent,omitempty"`
	InvertedTurn         []InvertedTurn         `xml:"inverted-turn,omitempty"`
	InvertedVerticalTurn []InvertedVerticalTurn `xml:"inverted-vertical-turn,omitempty"`
	Mordent              []Mordent              `xml:"mordent,omitempty"`
	OtherOrnament        []OtherOrnament        `xml:"other-ornament,omitempty"`
	Schleifer            []Schleifer            `xml:"schleifer,omitempty"`
	Shake                []Shake                `xml:"shake,omitempty"`
	Tremolo              []Tremolo              `xml:"tremolo,omitempty"`
	TrillMark            []TrillMark            `xml:"trill-mark,omitempty"`
	Turn                 []Turn                 `xml:"turn,omitempty"`
	VerticalTurn         []VerticalTurn         `xml:"vertical-turn,omitempty"`
	WavyLine             []WavyLine             `xml:"wavy-line,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Ornaments) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["id"] = r.IdAttr

	children["accidental-mark"] = make([]*MXML, len(r.AccidentalMark))
	for i, c := range r.AccidentalMark {
		children["accidental-mark"][i] = c.ToMXML()
	}

	children["delayed-inverted-turn"] = make([]*MXML, len(r.DelayedInvertedTurn))
	for i, c := range r.DelayedInvertedTurn {
		children["delayed-inverted-turn"][i] = c.ToMXML()
	}

	children["delayed-turn"] = make([]*MXML, len(r.DelayedTurn))
	for i, c := range r.DelayedTurn {
		children["delayed-turn"][i] = c.ToMXML()
	}

	children["haydn"] = make([]*MXML, len(r.Haydn))
	for i, c := range r.Haydn {
		children["haydn"][i] = c.ToMXML()
	}

	children["inverted-mordent"] = make([]*MXML, len(r.InvertedMordent))
	for i, c := range r.InvertedMordent {
		children["inverted-mordent"][i] = c.ToMXML()
	}

	children["inverted-turn"] = make([]*MXML, len(r.InvertedTurn))
	for i, c := range r.InvertedTurn {
		children["inverted-turn"][i] = c.ToMXML()
	}

	children["inverted-vertical-turn"] = make([]*MXML, len(r.InvertedVerticalTurn))
	for i, c := range r.InvertedVerticalTurn {
		children["inverted-vertical-turn"][i] = c.ToMXML()
	}

	children["mordent"] = make([]*MXML, len(r.Mordent))
	for i, c := range r.Mordent {
		children["mordent"][i] = c.ToMXML()
	}

	children["other-ornament"] = make([]*MXML, len(r.OtherOrnament))
	for i, c := range r.OtherOrnament {
		children["other-ornament"][i] = c.ToMXML()
	}

	children["schleifer"] = make([]*MXML, len(r.Schleifer))
	for i, c := range r.Schleifer {
		children["schleifer"][i] = c.ToMXML()
	}

	children["shake"] = make([]*MXML, len(r.Shake))
	for i, c := range r.Shake {
		children["shake"][i] = c.ToMXML()
	}

	children["tremolo"] = make([]*MXML, len(r.Tremolo))
	for i, c := range r.Tremolo {
		children["tremolo"][i] = c.ToMXML()
	}

	children["trill-mark"] = make([]*MXML, len(r.TrillMark))
	for i, c := range r.TrillMark {
		children["trill-mark"][i] = c.ToMXML()
	}

	children["turn"] = make([]*MXML, len(r.Turn))
	for i, c := range r.Turn {
		children["turn"][i] = c.ToMXML()
	}

	children["vertical-turn"] = make([]*MXML, len(r.VerticalTurn))
	for i, c := range r.VerticalTurn {
		children["vertical-turn"][i] = c.ToMXML()
	}

	children["wavy-line"] = make([]*MXML, len(r.WavyLine))
	for i, c := range r.WavyLine {
		children["wavy-line"][i] = c.ToMXML()
	}

	return newMXML("ornaments", strings.TrimSpace(r.IValue), attributes, children)
}
