package mxml

import (
	"encoding/xml"
	"strings"
)

// StaffTuning is the structure for staff-tuning MusicXML element.
type StaffTuning struct {
	XMLName xml.Name `xml:"staff-tuning"`

	LineAttr string `xml:"line,attr,omitempty"`

	TuningAlter  []TuningAlter  `xml:"tuning-alter,omitempty"`
	TuningOctave []TuningOctave `xml:"tuning-octave,omitempty"`
	TuningStep   []TuningStep   `xml:"tuning-step,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *StaffTuning) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["line"] = r.LineAttr

	children["tuning-alter"] = make([]*MXML, len(r.TuningAlter))
	for i, c := range r.TuningAlter {
		children["tuning-alter"][i] = c.ToMXML()
	}

	children["tuning-octave"] = make([]*MXML, len(r.TuningOctave))
	for i, c := range r.TuningOctave {
		children["tuning-octave"][i] = c.ToMXML()
	}

	children["tuning-step"] = make([]*MXML, len(r.TuningStep))
	for i, c := range r.TuningStep {
		children["tuning-step"][i] = c.ToMXML()
	}

	return newMXML("staff-tuning", strings.TrimSpace(r.IValue), attributes, children)
}
