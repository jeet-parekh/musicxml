package mxml

import (
	"encoding/xml"
	"strings"
)

// ScoreInstrument is the structure for score-instrument MusicXML element.
type ScoreInstrument struct {
	XMLName xml.Name `xml:"score-instrument"`

	IdAttr string `xml:"id,attr,omitempty"`

	Ensemble               []Ensemble               `xml:"ensemble,omitempty"`
	InstrumentAbbreviation []InstrumentAbbreviation `xml:"instrument-abbreviation,omitempty"`
	InstrumentName         []InstrumentName         `xml:"instrument-name,omitempty"`
	InstrumentSound        []InstrumentSound        `xml:"instrument-sound,omitempty"`
	Solo                   []Solo                   `xml:"solo,omitempty"`
	VirtualInstrument      []VirtualInstrument      `xml:"virtual-instrument,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *ScoreInstrument) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["id"] = r.IdAttr

	children["ensemble"] = make([]*MXML, len(r.Ensemble))
	for i, c := range r.Ensemble {
		children["ensemble"][i] = c.ToMXML()
	}

	children["instrument-abbreviation"] = make([]*MXML, len(r.InstrumentAbbreviation))
	for i, c := range r.InstrumentAbbreviation {
		children["instrument-abbreviation"][i] = c.ToMXML()
	}

	children["instrument-name"] = make([]*MXML, len(r.InstrumentName))
	for i, c := range r.InstrumentName {
		children["instrument-name"][i] = c.ToMXML()
	}

	children["instrument-sound"] = make([]*MXML, len(r.InstrumentSound))
	for i, c := range r.InstrumentSound {
		children["instrument-sound"][i] = c.ToMXML()
	}

	children["solo"] = make([]*MXML, len(r.Solo))
	for i, c := range r.Solo {
		children["solo"][i] = c.ToMXML()
	}

	children["virtual-instrument"] = make([]*MXML, len(r.VirtualInstrument))
	for i, c := range r.VirtualInstrument {
		children["virtual-instrument"][i] = c.ToMXML()
	}

	return newMXML("score-instrument", strings.TrimSpace(r.IValue), attributes, children)
}
