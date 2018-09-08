package mxml

import (
	"encoding/xml"
	"strings"
)

// Appearance is the structure for appearance MusicXML element.
type Appearance struct {
	XMLName xml.Name `xml:"appearance"`

	Distance        []Distance        `xml:"distance,omitempty"`
	Glyph           []Glyph           `xml:"glyph,omitempty"`
	LineWidth       []LineWidth       `xml:"line-width,omitempty"`
	NoteSize        []NoteSize        `xml:"note-size,omitempty"`
	OtherAppearance []OtherAppearance `xml:"other-appearance,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Appearance) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["distance"] = make([]*MXML, len(r.Distance))
	for i, c := range r.Distance {
		children["distance"][i] = c.ToMXML()
	}

	children["glyph"] = make([]*MXML, len(r.Glyph))
	for i, c := range r.Glyph {
		children["glyph"][i] = c.ToMXML()
	}

	children["line-width"] = make([]*MXML, len(r.LineWidth))
	for i, c := range r.LineWidth {
		children["line-width"][i] = c.ToMXML()
	}

	children["note-size"] = make([]*MXML, len(r.NoteSize))
	for i, c := range r.NoteSize {
		children["note-size"][i] = c.ToMXML()
	}

	children["other-appearance"] = make([]*MXML, len(r.OtherAppearance))
	for i, c := range r.OtherAppearance {
		children["other-appearance"][i] = c.ToMXML()
	}

	return newMXML("appearance", strings.TrimSpace(r.IValue), attributes, children)
}
