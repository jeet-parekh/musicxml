package mxml

import (
	"encoding/xml"
	"strings"
)

// Interchangeable is the structure for interchangeable MusicXML element.
type Interchangeable struct {
	XMLName xml.Name `xml:"interchangeable"`

	SeparatorAttr string `xml:"separator,attr,omitempty"`
	SymbolAttr    string `xml:"symbol,attr,omitempty"`

	BeatType     []BeatType     `xml:"beat-type,omitempty"`
	Beats        []Beats        `xml:"beats,omitempty"`
	TimeRelation []TimeRelation `xml:"time-relation,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Interchangeable) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["separator"] = r.SeparatorAttr
	attributes["symbol"] = r.SymbolAttr

	children["beat-type"] = make([]*MXML, len(r.BeatType))
	for i, c := range r.BeatType {
		children["beat-type"][i] = c.ToMXML()
	}

	children["beats"] = make([]*MXML, len(r.Beats))
	for i, c := range r.Beats {
		children["beats"][i] = c.ToMXML()
	}

	children["time-relation"] = make([]*MXML, len(r.TimeRelation))
	for i, c := range r.TimeRelation {
		children["time-relation"][i] = c.ToMXML()
	}

	return newMXML("interchangeable", strings.TrimSpace(r.IValue), attributes, children)
}
