package mxml

import (
	"encoding/xml"
	"strings"
)

// TupletActual is the structure for tuplet-actual MusicXML element.
type TupletActual struct {
	XMLName xml.Name `xml:"tuplet-actual"`

	TupletDot    []TupletDot    `xml:"tuplet-dot,omitempty"`
	TupletNumber []TupletNumber `xml:"tuplet-number,omitempty"`
	TupletType   []TupletType   `xml:"tuplet-type,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *TupletActual) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["tuplet-dot"] = make([]*MXML, len(r.TupletDot))
	for i, c := range r.TupletDot {
		children["tuplet-dot"][i] = c.ToMXML()
	}

	children["tuplet-number"] = make([]*MXML, len(r.TupletNumber))
	for i, c := range r.TupletNumber {
		children["tuplet-number"][i] = c.ToMXML()
	}

	children["tuplet-type"] = make([]*MXML, len(r.TupletType))
	for i, c := range r.TupletType {
		children["tuplet-type"][i] = c.ToMXML()
	}

	return newMXML("tuplet-actual", strings.TrimSpace(r.IValue), attributes, children)
}
