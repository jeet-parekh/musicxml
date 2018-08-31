package musicxml

import (
	"encoding/xml"
	"strings"
)

// TimeModification is the structure for time-modification MusicXML element.
type TimeModification struct {
	XMLName xml.Name `xml:"time-modification"`

	ActualNotes []ActualNotes `xml:"actual-notes,omitempty"`
	NormalDot   []NormalDot   `xml:"normal-dot,omitempty"`
	NormalNotes []NormalNotes `xml:"normal-notes,omitempty"`
	NormalType  []NormalType  `xml:"normal-type,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *TimeModification) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

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

	return newMXML("time-modification", strings.TrimSpace(r.IValue), attributes, children)
}
