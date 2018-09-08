package mxml

import (
	"encoding/xml"
	"strings"
)

// MeasureLayout is the structure for measure-layout MusicXML element.
type MeasureLayout struct {
	XMLName xml.Name `xml:"measure-layout"`

	MeasureDistance []MeasureDistance `xml:"measure-distance,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MeasureLayout) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["measure-distance"] = make([]*MXML, len(r.MeasureDistance))
	for i, c := range r.MeasureDistance {
		children["measure-distance"][i] = c.ToMXML()
	}

	return newMXML("measure-layout", strings.TrimSpace(r.IValue), attributes, children)
}
