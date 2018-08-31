package musicxml

import (
	"encoding/xml"
	"strings"
)

// Transpose is the structure for transpose MusicXML element.
type Transpose struct {
	XMLName xml.Name `xml:"transpose"`

	IdAttr     string `xml:"id,attr,omitempty"`
	NumberAttr string `xml:"number,attr,omitempty"`

	Chromatic    []Chromatic    `xml:"chromatic,omitempty"`
	Diatonic     []Diatonic     `xml:"diatonic,omitempty"`
	Double       []Double       `xml:"double,omitempty"`
	OctaveChange []OctaveChange `xml:"octave-change,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Transpose) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["id"] = r.IdAttr
	attributes["number"] = r.NumberAttr

	children["chromatic"] = make([]*MXML, len(r.Chromatic))
	for i, c := range r.Chromatic {
		children["chromatic"][i] = c.ToMXML()
	}

	children["diatonic"] = make([]*MXML, len(r.Diatonic))
	for i, c := range r.Diatonic {
		children["diatonic"][i] = c.ToMXML()
	}

	children["double"] = make([]*MXML, len(r.Double))
	for i, c := range r.Double {
		children["double"][i] = c.ToMXML()
	}

	children["octave-change"] = make([]*MXML, len(r.OctaveChange))
	for i, c := range r.OctaveChange {
		children["octave-change"][i] = c.ToMXML()
	}

	return newMXML("transpose", strings.TrimSpace(r.IValue), attributes, children)
}
