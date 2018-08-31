package musicxml

import (
	"encoding/xml"
	"strings"
)

// MetronomeTuplet is the structure for metronome-tuplet MusicXML element.
type MetronomeTuplet struct {
	XMLName xml.Name `xml:"metronome-tuplet"`

	BracketAttr    string `xml:"bracket,attr,omitempty"`
	ShowNumberAttr string `xml:"show-number,attr,omitempty"`
	TypeAttr       string `xml:"type,attr,omitempty"`

	ActualNotes []ActualNotes `xml:"actual-notes,omitempty"`
	NormalDot   []NormalDot   `xml:"normal-dot,omitempty"`
	NormalNotes []NormalNotes `xml:"normal-notes,omitempty"`
	NormalType  []NormalType  `xml:"normal-type,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MetronomeTuplet) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["bracket"] = r.BracketAttr
	attributes["show-number"] = r.ShowNumberAttr
	attributes["type"] = r.TypeAttr

	children["actual-notes"] = make([]*MXML, len(r.ActualNotes))
	for i, c := range r.ActualNotes {
		children["actual-notes"][i] = c.ToMXML()
	}

	children["normal-dot"] = make([]*MXML, len(r.NormalDot))
	for i, c := range r.NormalDot {
		children["normal-dot"][i] = c.ToMXML()
	}

	children["normal-notes"] = make([]*MXML, len(r.NormalNotes))
	for i, c := range r.NormalNotes {
		children["normal-notes"][i] = c.ToMXML()
	}

	children["normal-type"] = make([]*MXML, len(r.NormalType))
	for i, c := range r.NormalType {
		children["normal-type"][i] = c.ToMXML()
	}

	return newMXML("metronome-tuplet", strings.TrimSpace(r.IValue), attributes, children)
}
