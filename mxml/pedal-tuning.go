package mxml

import (
	"encoding/xml"
	"strings"
)

// PedalTuning is the structure for pedal-tuning MusicXML element.
type PedalTuning struct {
	XMLName xml.Name `xml:"pedal-tuning"`

	PedalAlter []PedalAlter `xml:"pedal-alter,omitempty"`
	PedalStep  []PedalStep  `xml:"pedal-step,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *PedalTuning) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["pedal-alter"] = make([]*MXML, len(r.PedalAlter))
	for i, c := range r.PedalAlter {
		children["pedal-alter"][i] = c.ToMXML()
	}

	children["pedal-step"] = make([]*MXML, len(r.PedalStep))
	for i, c := range r.PedalStep {
		children["pedal-step"][i] = c.ToMXML()
	}

	return newMXML("pedal-tuning", strings.TrimSpace(r.IValue), attributes, children)
}
