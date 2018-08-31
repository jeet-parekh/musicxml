package musicxml

import (
	"encoding/xml"
	"strings"
)

// Work is the structure for work MusicXML element.
type Work struct {
	XMLName xml.Name `xml:"work"`

	Opus       []Opus       `xml:"opus,omitempty"`
	WorkNumber []WorkNumber `xml:"work-number,omitempty"`
	WorkTitle  []WorkTitle  `xml:"work-title,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Work) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["opus"] = make([]*MXML, len(r.Opus))
	for i, c := range r.Opus {
		children["opus"][i] = c.ToMXML()
	}

	children["work-number"] = make([]*MXML, len(r.WorkNumber))
	for i, c := range r.WorkNumber {
		children["work-number"][i] = c.ToMXML()
	}

	children["work-title"] = make([]*MXML, len(r.WorkTitle))
	for i, c := range r.WorkTitle {
		children["work-title"][i] = c.ToMXML()
	}

	return newMXML("work", strings.TrimSpace(r.IValue), attributes, children)
}
