package mxml

import (
	"encoding/xml"
	"strings"
)

// MetronomeNote is the structure for metronome-note MusicXML element.
type MetronomeNote struct {
	XMLName xml.Name `xml:"metronome-note"`

	MetronomeBeam   []MetronomeBeam   `xml:"metronome-beam,omitempty"`
	MetronomeDot    []MetronomeDot    `xml:"metronome-dot,omitempty"`
	MetronomeTied   []MetronomeTied   `xml:"metronome-tied,omitempty"`
	MetronomeTuplet []MetronomeTuplet `xml:"metronome-tuplet,omitempty"`
	MetronomeType   []MetronomeType   `xml:"metronome-type,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MetronomeNote) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["metronome-beam"] = make([]*MXML, len(r.MetronomeBeam))
	for i, c := range r.MetronomeBeam {
		children["metronome-beam"][i] = c.ToMXML()
	}

	children["metronome-dot"] = make([]*MXML, len(r.MetronomeDot))
	for i, c := range r.MetronomeDot {
		children["metronome-dot"][i] = c.ToMXML()
	}

	children["metronome-tied"] = make([]*MXML, len(r.MetronomeTied))
	for i, c := range r.MetronomeTied {
		children["metronome-tied"][i] = c.ToMXML()
	}

	children["metronome-tuplet"] = make([]*MXML, len(r.MetronomeTuplet))
	for i, c := range r.MetronomeTuplet {
		children["metronome-tuplet"][i] = c.ToMXML()
	}

	children["metronome-type"] = make([]*MXML, len(r.MetronomeType))
	for i, c := range r.MetronomeType {
		children["metronome-type"][i] = c.ToMXML()
	}

	return newMXML("metronome-note", strings.TrimSpace(r.IValue), attributes, children)
}
