package mxml

import (
	"encoding/xml"
	"strings"
)

// Articulations is the structure for articulations MusicXML element.
type Articulations struct {
	XMLName xml.Name `xml:"articulations"`

	IdAttr string `xml:"id,attr,omitempty"`

	Accent            []Accent            `xml:"accent,omitempty"`
	BreathMark        []BreathMark        `xml:"breath-mark,omitempty"`
	Caesura           []Caesura           `xml:"caesura,omitempty"`
	DetachedLegato    []DetachedLegato    `xml:"detached-legato,omitempty"`
	Doit              []Doit              `xml:"doit,omitempty"`
	Falloff           []Falloff           `xml:"falloff,omitempty"`
	OtherArticulation []OtherArticulation `xml:"other-articulation,omitempty"`
	Plop              []Plop              `xml:"plop,omitempty"`
	Scoop             []Scoop             `xml:"scoop,omitempty"`
	SoftAccent        []SoftAccent        `xml:"soft-accent,omitempty"`
	Spiccato          []Spiccato          `xml:"spiccato,omitempty"`
	Staccatissimo     []Staccatissimo     `xml:"staccatissimo,omitempty"`
	Staccato          []Staccato          `xml:"staccato,omitempty"`
	Stress            []Stress            `xml:"stress,omitempty"`
	StrongAccent      []StrongAccent      `xml:"strong-accent,omitempty"`
	Tenuto            []Tenuto            `xml:"tenuto,omitempty"`
	Unstress          []Unstress          `xml:"unstress,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Articulations) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["id"] = r.IdAttr

	children["accent"] = make([]*MXML, len(r.Accent))
	for i, c := range r.Accent {
		children["accent"][i] = c.ToMXML()
	}

	children["breath-mark"] = make([]*MXML, len(r.BreathMark))
	for i, c := range r.BreathMark {
		children["breath-mark"][i] = c.ToMXML()
	}

	children["caesura"] = make([]*MXML, len(r.Caesura))
	for i, c := range r.Caesura {
		children["caesura"][i] = c.ToMXML()
	}

	children["detached-legato"] = make([]*MXML, len(r.DetachedLegato))
	for i, c := range r.DetachedLegato {
		children["detached-legato"][i] = c.ToMXML()
	}

	children["doit"] = make([]*MXML, len(r.Doit))
	for i, c := range r.Doit {
		children["doit"][i] = c.ToMXML()
	}

	children["falloff"] = make([]*MXML, len(r.Falloff))
	for i, c := range r.Falloff {
		children["falloff"][i] = c.ToMXML()
	}

	children["other-articulation"] = make([]*MXML, len(r.OtherArticulation))
	for i, c := range r.OtherArticulation {
		children["other-articulation"][i] = c.ToMXML()
	}

	children["plop"] = make([]*MXML, len(r.Plop))
	for i, c := range r.Plop {
		children["plop"][i] = c.ToMXML()
	}

	children["scoop"] = make([]*MXML, len(r.Scoop))
	for i, c := range r.Scoop {
		children["scoop"][i] = c.ToMXML()
	}

	children["soft-accent"] = make([]*MXML, len(r.SoftAccent))
	for i, c := range r.SoftAccent {
		children["soft-accent"][i] = c.ToMXML()
	}

	children["spiccato"] = make([]*MXML, len(r.Spiccato))
	for i, c := range r.Spiccato {
		children["spiccato"][i] = c.ToMXML()
	}

	children["staccatissimo"] = make([]*MXML, len(r.Staccatissimo))
	for i, c := range r.Staccatissimo {
		children["staccatissimo"][i] = c.ToMXML()
	}

	children["staccato"] = make([]*MXML, len(r.Staccato))
	for i, c := range r.Staccato {
		children["staccato"][i] = c.ToMXML()
	}

	children["stress"] = make([]*MXML, len(r.Stress))
	for i, c := range r.Stress {
		children["stress"][i] = c.ToMXML()
	}

	children["strong-accent"] = make([]*MXML, len(r.StrongAccent))
	for i, c := range r.StrongAccent {
		children["strong-accent"][i] = c.ToMXML()
	}

	children["tenuto"] = make([]*MXML, len(r.Tenuto))
	for i, c := range r.Tenuto {
		children["tenuto"][i] = c.ToMXML()
	}

	children["unstress"] = make([]*MXML, len(r.Unstress))
	for i, c := range r.Unstress {
		children["unstress"][i] = c.ToMXML()
	}

	return newMXML("articulations", strings.TrimSpace(r.IValue), attributes, children)
}
