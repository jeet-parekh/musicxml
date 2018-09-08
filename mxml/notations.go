package mxml

import (
	"encoding/xml"
	"strings"
)

// Notations is the structure for notations MusicXML element.
type Notations struct {
	XMLName xml.Name `xml:"notations"`

	IdAttr          string `xml:"id,attr,omitempty"`
	PrintObjectAttr string `xml:"print-object,attr,omitempty"`

	AccidentalMark []AccidentalMark `xml:"accidental-mark,omitempty"`
	Arpeggiate     []Arpeggiate     `xml:"arpeggiate,omitempty"`
	Articulations  []Articulations  `xml:"articulations,omitempty"`
	Dynamics       []Dynamics       `xml:"dynamics,omitempty"`
	Fermata        []Fermata        `xml:"fermata,omitempty"`
	Footnote       []Footnote       `xml:"footnote,omitempty"`
	Glissando      []Glissando      `xml:"glissando,omitempty"`
	Level          []Level          `xml:"level,omitempty"`
	NonArpeggiate  []NonArpeggiate  `xml:"non-arpeggiate,omitempty"`
	Ornaments      []Ornaments      `xml:"ornaments,omitempty"`
	OtherNotation  []OtherNotation  `xml:"other-notation,omitempty"`
	Slide          []Slide          `xml:"slide,omitempty"`
	Slur           []Slur           `xml:"slur,omitempty"`
	Technical      []Technical      `xml:"technical,omitempty"`
	Tied           []Tied           `xml:"tied,omitempty"`
	Tuplet         []Tuplet         `xml:"tuplet,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Notations) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["id"] = r.IdAttr
	attributes["print-object"] = r.PrintObjectAttr

	children["accidental-mark"] = make([]*MXML, len(r.AccidentalMark))
	for i, c := range r.AccidentalMark {
		children["accidental-mark"][i] = c.ToMXML()
	}

	children["arpeggiate"] = make([]*MXML, len(r.Arpeggiate))
	for i, c := range r.Arpeggiate {
		children["arpeggiate"][i] = c.ToMXML()
	}

	children["articulations"] = make([]*MXML, len(r.Articulations))
	for i, c := range r.Articulations {
		children["articulations"][i] = c.ToMXML()
	}

	children["dynamics"] = make([]*MXML, len(r.Dynamics))
	for i, c := range r.Dynamics {
		children["dynamics"][i] = c.ToMXML()
	}

	children["fermata"] = make([]*MXML, len(r.Fermata))
	for i, c := range r.Fermata {
		children["fermata"][i] = c.ToMXML()
	}

	children["footnote"] = make([]*MXML, len(r.Footnote))
	for i, c := range r.Footnote {
		children["footnote"][i] = c.ToMXML()
	}

	children["glissando"] = make([]*MXML, len(r.Glissando))
	for i, c := range r.Glissando {
		children["glissando"][i] = c.ToMXML()
	}

	children["level"] = make([]*MXML, len(r.Level))
	for i, c := range r.Level {
		children["level"][i] = c.ToMXML()
	}

	children["non-arpeggiate"] = make([]*MXML, len(r.NonArpeggiate))
	for i, c := range r.NonArpeggiate {
		children["non-arpeggiate"][i] = c.ToMXML()
	}

	children["ornaments"] = make([]*MXML, len(r.Ornaments))
	for i, c := range r.Ornaments {
		children["ornaments"][i] = c.ToMXML()
	}

	children["other-notation"] = make([]*MXML, len(r.OtherNotation))
	for i, c := range r.OtherNotation {
		children["other-notation"][i] = c.ToMXML()
	}

	children["slide"] = make([]*MXML, len(r.Slide))
	for i, c := range r.Slide {
		children["slide"][i] = c.ToMXML()
	}

	children["slur"] = make([]*MXML, len(r.Slur))
	for i, c := range r.Slur {
		children["slur"][i] = c.ToMXML()
	}

	children["technical"] = make([]*MXML, len(r.Technical))
	for i, c := range r.Technical {
		children["technical"][i] = c.ToMXML()
	}

	children["tied"] = make([]*MXML, len(r.Tied))
	for i, c := range r.Tied {
		children["tied"][i] = c.ToMXML()
	}

	children["tuplet"] = make([]*MXML, len(r.Tuplet))
	for i, c := range r.Tuplet {
		children["tuplet"][i] = c.ToMXML()
	}

	return newMXML("notations", strings.TrimSpace(r.IValue), attributes, children)
}
