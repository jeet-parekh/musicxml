package mxml

import (
	"encoding/xml"
	"strings"
)

// MidiInstrument is the structure for midi-instrument MusicXML element.
type MidiInstrument struct {
	XMLName xml.Name `xml:"midi-instrument"`

	IdAttr string `xml:"id,attr,omitempty"`

	Elevation     []Elevation     `xml:"elevation,omitempty"`
	MidiBank      []MidiBank      `xml:"midi-bank,omitempty"`
	MidiChannel   []MidiChannel   `xml:"midi-channel,omitempty"`
	MidiName      []MidiName      `xml:"midi-name,omitempty"`
	MidiProgram   []MidiProgram   `xml:"midi-program,omitempty"`
	MidiUnpitched []MidiUnpitched `xml:"midi-unpitched,omitempty"`
	Pan           []Pan           `xml:"pan,omitempty"`
	Volume        []Volume        `xml:"volume,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MidiInstrument) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["id"] = r.IdAttr

	children["elevation"] = make([]*MXML, len(r.Elevation))
	for i, c := range r.Elevation {
		children["elevation"][i] = c.ToMXML()
	}

	children["midi-bank"] = make([]*MXML, len(r.MidiBank))
	for i, c := range r.MidiBank {
		children["midi-bank"][i] = c.ToMXML()
	}

	children["midi-channel"] = make([]*MXML, len(r.MidiChannel))
	for i, c := range r.MidiChannel {
		children["midi-channel"][i] = c.ToMXML()
	}

	children["midi-name"] = make([]*MXML, len(r.MidiName))
	for i, c := range r.MidiName {
		children["midi-name"][i] = c.ToMXML()
	}

	children["midi-program"] = make([]*MXML, len(r.MidiProgram))
	for i, c := range r.MidiProgram {
		children["midi-program"][i] = c.ToMXML()
	}

	children["midi-unpitched"] = make([]*MXML, len(r.MidiUnpitched))
	for i, c := range r.MidiUnpitched {
		children["midi-unpitched"][i] = c.ToMXML()
	}

	children["pan"] = make([]*MXML, len(r.Pan))
	for i, c := range r.Pan {
		children["pan"][i] = c.ToMXML()
	}

	children["volume"] = make([]*MXML, len(r.Volume))
	for i, c := range r.Volume {
		children["volume"][i] = c.ToMXML()
	}

	return newMXML("midi-instrument", strings.TrimSpace(r.IValue), attributes, children)
}
