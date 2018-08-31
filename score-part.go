package musicxml

import (
	"encoding/xml"
	"strings"
)

// ScorePart is the structure for score-part MusicXML element.
type ScorePart struct {
	XMLName xml.Name `xml:"score-part"`

	IdAttr string `xml:"id,attr,omitempty"`

	Group                   []Group                   `xml:"group,omitempty"`
	Identification          []Identification          `xml:"identification,omitempty"`
	MidiDevice              []MidiDevice              `xml:"midi-device,omitempty"`
	MidiInstrument          []MidiInstrument          `xml:"midi-instrument,omitempty"`
	PartAbbreviation        []PartAbbreviation        `xml:"part-abbreviation,omitempty"`
	PartAbbreviationDisplay []PartAbbreviationDisplay `xml:"part-abbreviation-display,omitempty"`
	PartName                []PartName                `xml:"part-name,omitempty"`
	PartNameDisplay         []PartNameDisplay         `xml:"part-name-display,omitempty"`
	ScoreInstrument         []ScoreInstrument         `xml:"score-instrument,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *ScorePart) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["id"] = r.IdAttr

	children["group"] = make([]*MXML, len(r.Group))
	for i, c := range r.Group {
		children["group"][i] = c.ToMXML()
	}

	children["identification"] = make([]*MXML, len(r.Identification))
	for i, c := range r.Identification {
		children["identification"][i] = c.ToMXML()
	}

	children["midi-device"] = make([]*MXML, len(r.MidiDevice))
	for i, c := range r.MidiDevice {
		children["midi-device"][i] = c.ToMXML()
	}

	children["midi-instrument"] = make([]*MXML, len(r.MidiInstrument))
	for i, c := range r.MidiInstrument {
		children["midi-instrument"][i] = c.ToMXML()
	}

	children["part-abbreviation"] = make([]*MXML, len(r.PartAbbreviation))
	for i, c := range r.PartAbbreviation {
		children["part-abbreviation"][i] = c.ToMXML()
	}

	children["part-abbreviation-display"] = make([]*MXML, len(r.PartAbbreviationDisplay))
	for i, c := range r.PartAbbreviationDisplay {
		children["part-abbreviation-display"][i] = c.ToMXML()
	}

	children["part-name"] = make([]*MXML, len(r.PartName))
	for i, c := range r.PartName {
		children["part-name"][i] = c.ToMXML()
	}

	children["part-name-display"] = make([]*MXML, len(r.PartNameDisplay))
	for i, c := range r.PartNameDisplay {
		children["part-name-display"][i] = c.ToMXML()
	}

	children["score-instrument"] = make([]*MXML, len(r.ScoreInstrument))
	for i, c := range r.ScoreInstrument {
		children["score-instrument"][i] = c.ToMXML()
	}

	return newMXML("score-part", strings.TrimSpace(r.IValue), attributes, children)
}
