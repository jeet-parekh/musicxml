package musicxml

import (
	"encoding/xml"
	"strings"
)

// GroupAbbreviationDisplay is the structure for group-abbreviation-display MusicXML element.
type GroupAbbreviationDisplay struct {
	XMLName xml.Name `xml:"group-abbreviation-display"`

	PrintObjectAttr string `xml:"print-object,attr,omitempty"`

	AccidentalText []AccidentalText `xml:"accidental-text,omitempty"`
	DisplayText    []DisplayText    `xml:"display-text,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *GroupAbbreviationDisplay) ToMXML() *MXML {
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

	return newMXML("group-abbreviation-display", strings.TrimSpace(r.IValue), attributes, children)
}
