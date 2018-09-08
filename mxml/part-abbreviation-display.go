package mxml

import (
	"encoding/xml"
	"strings"
)

// PartAbbreviationDisplay is the structure for part-abbreviation-display MusicXML element.
type PartAbbreviationDisplay struct {
	XMLName xml.Name `xml:"part-abbreviation-display"`

	PrintObjectAttr string `xml:"print-object,attr,omitempty"`

	AccidentalText []AccidentalText `xml:"accidental-text,omitempty"`
	DisplayText    []DisplayText    `xml:"display-text,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *PartAbbreviationDisplay) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["print-object"] = r.PrintObjectAttr

	children["accidental-text"] = make([]*MXML, len(r.AccidentalText))
	for i, c := range r.AccidentalText {
		children["accidental-text"][i] = c.ToMXML()
	}

	children["display-text"] = make([]*MXML, len(r.DisplayText))
	for i, c := range r.DisplayText {
		children["display-text"][i] = c.ToMXML()
	}

	return newMXML("part-abbreviation-display", strings.TrimSpace(r.IValue), attributes, children)
}
