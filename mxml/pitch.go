package mxml

import (
	"encoding/xml"
	"strings"
)

// Pitch is the structure for pitch MusicXML element.
type Pitch struct {
	XMLName xml.Name `xml:"pitch"`

	Alter  []Alter  `xml:"alter,omitempty"`
	Octave []Octave `xml:"octave,omitempty"`
	Step   []Step   `xml:"step,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Pitch) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["alter"] = make([]*MXML, len(r.Alter))
	for i, c := range r.Alter {
		children["alter"][i] = c.ToMXML()
	}

	children["octave"] = make([]*MXML, len(r.Octave))
	for i, c := range r.Octave {
		children["octave"][i] = c.ToMXML()
	}

	children["step"] = make([]*MXML, len(r.Step))
	for i, c := range r.Step {
		children["step"][i] = c.ToMXML()
	}

	return newMXML("pitch", strings.TrimSpace(r.IValue), attributes, children)
}
