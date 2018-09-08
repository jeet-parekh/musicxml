package mxml

import (
	"encoding/xml"
	"strings"
)

// GroupNameDisplay is the structure for group-name-display MusicXML element.
type GroupNameDisplay struct {
	XMLName xml.Name `xml:"group-name-display"`

	PrintObjectAttr string `xml:"print-object,attr,omitempty"`

	AccidentalText []AccidentalText `xml:"accidental-text,omitempty"`
	DisplayText    []DisplayText    `xml:"display-text,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *GroupNameDisplay) ToMXML() *MXML {
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

	return newMXML("group-name-display", strings.TrimSpace(r.IValue), attributes, children)
}
