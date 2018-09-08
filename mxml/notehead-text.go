package mxml

import (
	"encoding/xml"
	"strings"
)

// NoteheadText is the structure for notehead-text MusicXML element.
type NoteheadText struct {
	XMLName xml.Name `xml:"notehead-text"`

	AccidentalText []AccidentalText `xml:"accidental-text,omitempty"`
	DisplayText    []DisplayText    `xml:"display-text,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *NoteheadText) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["accidental-text"] = make([]*MXML, len(r.AccidentalText))
	for i, c := range r.AccidentalText {
		children["accidental-text"][i] = c.ToMXML()
	}

	children["display-text"] = make([]*MXML, len(r.DisplayText))
	for i, c := range r.DisplayText {
		children["display-text"][i] = c.ToMXML()
	}

	return newMXML("notehead-text", strings.TrimSpace(r.IValue), attributes, children)
}
