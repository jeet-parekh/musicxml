package mxml

import (
	"encoding/xml"
	"strings"
)

// Sound is the structure for sound MusicXML element.
type Sound struct {
	XMLName xml.Name `xml:"sound"`

	CodaAttr           string `xml:"coda,attr,omitempty"`
	DacapoAttr         string `xml:"dacapo,attr,omitempty"`
	DalsegnoAttr       string `xml:"dalsegno,attr,omitempty"`
	DamperPedalAttr    string `xml:"damper-pedal,attr,omitempty"`
	DivisionsAttr      string `xml:"divisions,attr,omitempty"`
	DynamicsAttr       string `xml:"dynamics,attr,omitempty"`
	ElevationAttr      string `xml:"elevation,attr,omitempty"`
	FineAttr           string `xml:"fine,attr,omitempty"`
	ForwardRepeatAttr  string `xml:"forward-repeat,attr,omitempty"`
	IdAttr             string `xml:"id,attr,omitempty"`
	PanAttr            string `xml:"pan,attr,omitempty"`
	PizzicatoAttr      string `xml:"pizzicato,attr,omitempty"`
	SegnoAttr          string `xml:"segno,attr,omitempty"`
	SoftPedalAttr      string `xml:"soft-pedal,attr,omitempty"`
	SostenutoPedalAttr string `xml:"sostenuto-pedal,attr,omitempty"`
	TempoAttr          string `xml:"tempo,attr,omitempty"`
	TimeOnlyAttr       string `xml:"time-only,attr,omitempty"`
	TocodaAttr         string `xml:"tocoda,attr,omitempty"`

	MidiDevice     []MidiDevice     `xml:"midi-device,omitempty"`
	MidiInstrument []MidiInstrument `xml:"midi-instrument,omitempty"`
	Offset         []Offset         `xml:"offset,omitempty"`
	Play           []Play           `xml:"play,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Sound) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["coda"] = r.CodaAttr
	attributes["dacapo"] = r.DacapoAttr
	attributes["dalsegno"] = r.DalsegnoAttr
	attributes["damper-pedal"] = r.DamperPedalAttr
	attributes["divisions"] = r.DivisionsAttr
	attributes["dynamics"] = r.DynamicsAttr
	attributes["elevation"] = r.ElevationAttr
	attributes["fine"] = r.FineAttr
	attributes["forward-repeat"] = r.ForwardRepeatAttr
	attributes["id"] = r.IdAttr
	attributes["pan"] = r.PanAttr
	attributes["pizzicato"] = r.PizzicatoAttr
	attributes["segno"] = r.SegnoAttr
	attributes["soft-pedal"] = r.SoftPedalAttr
	attributes["sostenuto-pedal"] = r.SostenutoPedalAttr
	attributes["tempo"] = r.TempoAttr
	attributes["time-only"] = r.TimeOnlyAttr
	attributes["tocoda"] = r.TocodaAttr

	children["midi-device"] = make([]*MXML, len(r.MidiDevice))
	for i, c := range r.MidiDevice {
		children["midi-device"][i] = c.ToMXML()
	}

	children["midi-instrument"] = make([]*MXML, len(r.MidiInstrument))
	for i, c := range r.MidiInstrument {
		children["midi-instrument"][i] = c.ToMXML()
	}

	children["offset"] = make([]*MXML, len(r.Offset))
	for i, c := range r.Offset {
		children["offset"][i] = c.ToMXML()
	}

	children["play"] = make([]*MXML, len(r.Play))
	for i, c := range r.Play {
		children["play"][i] = c.ToMXML()
	}

	return newMXML("sound", strings.TrimSpace(r.IValue), attributes, children)
}
