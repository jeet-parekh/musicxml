package mxml

import (
	"encoding/xml"
	"strings"
)

// Barline is the structure for barline MusicXML element.
type Barline struct {
	XMLName xml.Name `xml:"barline"`

	CodaAttr      string `xml:"coda,attr,omitempty"`
	DivisionsAttr string `xml:"divisions,attr,omitempty"`
	IdAttr        string `xml:"id,attr,omitempty"`
	LocationAttr  string `xml:"location,attr,omitempty"`
	SegnoAttr     string `xml:"segno,attr,omitempty"`

	BarStyle []BarStyle `xml:"bar-style,omitempty"`
	Coda     []Coda     `xml:"coda,omitempty"`
	Ending   []Ending   `xml:"ending,omitempty"`
	Fermata  []Fermata  `xml:"fermata,omitempty"`
	Footnote []Footnote `xml:"footnote,omitempty"`
	Level    []Level    `xml:"level,omitempty"`
	Repeat   []Repeat   `xml:"repeat,omitempty"`
	Segno    []Segno    `xml:"segno,omitempty"`
	WavyLine []WavyLine `xml:"wavy-line,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Barline) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["coda"] = r.CodaAttr
	attributes["divisions"] = r.DivisionsAttr
	attributes["id"] = r.IdAttr
	attributes["location"] = r.LocationAttr
	attributes["segno"] = r.SegnoAttr

	children["bar-style"] = make([]*MXML, len(r.BarStyle))
	for i, c := range r.BarStyle {
		children["bar-style"][i] = c.ToMXML()
	}

	children["coda"] = make([]*MXML, len(r.Coda))
	for i, c := range r.Coda {
		children["coda"][i] = c.ToMXML()
	}

	children["ending"] = make([]*MXML, len(r.Ending))
	for i, c := range r.Ending {
		children["ending"][i] = c.ToMXML()
	}

	children["fermata"] = make([]*MXML, len(r.Fermata))
	for i, c := range r.Fermata {
		children["fermata"][i] = c.ToMXML()
	}

	children["footnote"] = make([]*MXML, len(r.Footnote))
	for i, c := range r.Footnote {
		children["footnote"][i] = c.ToMXML()
	}

	children["level"] = make([]*MXML, len(r.Level))
	for i, c := range r.Level {
		children["level"][i] = c.ToMXML()
	}

	children["repeat"] = make([]*MXML, len(r.Repeat))
	for i, c := range r.Repeat {
		children["repeat"][i] = c.ToMXML()
	}

	children["segno"] = make([]*MXML, len(r.Segno))
	for i, c := range r.Segno {
		children["segno"][i] = c.ToMXML()
	}

	children["wavy-line"] = make([]*MXML, len(r.WavyLine))
	for i, c := range r.WavyLine {
		children["wavy-line"][i] = c.ToMXML()
	}

	return newMXML("barline", strings.TrimSpace(r.IValue), attributes, children)
}
