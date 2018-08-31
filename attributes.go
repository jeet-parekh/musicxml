package musicxml

import (
	"encoding/xml"
	"strings"
)

// Attributes is the structure for attributes MusicXML element.
type Attributes struct {
	XMLName xml.Name `xml:"attributes"`

	Clef         []Clef         `xml:"clef,omitempty"`
	Directive    []Directive    `xml:"directive,omitempty"`
	Divisions    []Divisions    `xml:"divisions,omitempty"`
	Footnote     []Footnote     `xml:"footnote,omitempty"`
	Instruments  []Instruments  `xml:"instruments,omitempty"`
	Key          []Key          `xml:"key,omitempty"`
	Level        []Level        `xml:"level,omitempty"`
	MeasureStyle []MeasureStyle `xml:"measure-style,omitempty"`
	PartSymbol   []PartSymbol   `xml:"part-symbol,omitempty"`
	StaffDetails []StaffDetails `xml:"staff-details,omitempty"`
	Staves       []Staves       `xml:"staves,omitempty"`
	Time         []Time         `xml:"time,omitempty"`
	Transpose    []Transpose    `xml:"transpose,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Attributes) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["clef"] = make([]*MXML, len(r.Clef))
	for i, c := range r.Clef {
		children["clef"][i] = c.ToMXML()
	}

	children["directive"] = make([]*MXML, len(r.Directive))
	for i, c := range r.Directive {
		children["directive"][i] = c.ToMXML()
	}

	children["divisions"] = make([]*MXML, len(r.Divisions))
	for i, c := range r.Divisions {
		children["divisions"][i] = c.ToMXML()
	}

	children["footnote"] = make([]*MXML, len(r.Footnote))
	for i, c := range r.Footnote {
		children["footnote"][i] = c.ToMXML()
	}

	children["instruments"] = make([]*MXML, len(r.Instruments))
	for i, c := range r.Instruments {
		children["instruments"][i] = c.ToMXML()
	}

	children["key"] = make([]*MXML, len(r.Key))
	for i, c := range r.Key {
		children["key"][i] = c.ToMXML()
	}

	children["level"] = make([]*MXML, len(r.Level))
	for i, c := range r.Level {
		children["level"][i] = c.ToMXML()
	}

	children["measure-style"] = make([]*MXML, len(r.MeasureStyle))
	for i, c := range r.MeasureStyle {
		children["measure-style"][i] = c.ToMXML()
	}

	children["part-symbol"] = make([]*MXML, len(r.PartSymbol))
	for i, c := range r.PartSymbol {
		children["part-symbol"][i] = c.ToMXML()
	}

	children["staff-details"] = make([]*MXML, len(r.StaffDetails))
	for i, c := range r.StaffDetails {
		children["staff-details"][i] = c.ToMXML()
	}

	children["staves"] = make([]*MXML, len(r.Staves))
	for i, c := range r.Staves {
		children["staves"][i] = c.ToMXML()
	}

	children["time"] = make([]*MXML, len(r.Time))
	for i, c := range r.Time {
		children["time"][i] = c.ToMXML()
	}

	children["transpose"] = make([]*MXML, len(r.Transpose))
	for i, c := range r.Transpose {
		children["transpose"][i] = c.ToMXML()
	}

	return newMXML("attributes", strings.TrimSpace(r.IValue), attributes, children)
}
