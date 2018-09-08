package mxml

import (
	"encoding/xml"
	"strings"
)

// Figure is the structure for figure MusicXML element.
type Figure struct {
	XMLName xml.Name `xml:"figure"`

	Extend       []Extend       `xml:"extend,omitempty"`
	FigureNumber []FigureNumber `xml:"figure-number,omitempty"`
	Footnote     []Footnote     `xml:"footnote,omitempty"`
	Level        []Level        `xml:"level,omitempty"`
	Prefix       []Prefix       `xml:"prefix,omitempty"`
	Suffix       []Suffix       `xml:"suffix,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Figure) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["extend"] = make([]*MXML, len(r.Extend))
	for i, c := range r.Extend {
		children["extend"][i] = c.ToMXML()
	}

	children["figure-number"] = make([]*MXML, len(r.FigureNumber))
	for i, c := range r.FigureNumber {
		children["figure-number"][i] = c.ToMXML()
	}

	children["footnote"] = make([]*MXML, len(r.Footnote))
	for i, c := range r.Footnote {
		children["footnote"][i] = c.ToMXML()
	}

	children["level"] = make([]*MXML, len(r.Level))
	for i, c := range r.Level {
		children["level"][i] = c.ToMXML()
	}

	children["prefix"] = make([]*MXML, len(r.Prefix))
	for i, c := range r.Prefix {
		children["prefix"][i] = c.ToMXML()
	}

	children["suffix"] = make([]*MXML, len(r.Suffix))
	for i, c := range r.Suffix {
		children["suffix"][i] = c.ToMXML()
	}

	return newMXML("figure", strings.TrimSpace(r.IValue), attributes, children)
}
