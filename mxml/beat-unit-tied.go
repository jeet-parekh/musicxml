package mxml

import (
	"encoding/xml"
	"strings"
)

// BeatUnitTied is the structure for beat-unit-tied MusicXML element.
type BeatUnitTied struct {
	XMLName xml.Name `xml:"beat-unit-tied"`

	BeatUnit    []BeatUnit    `xml:"beat-unit,omitempty"`
	BeatUnitDot []BeatUnitDot `xml:"beat-unit-dot,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *BeatUnitTied) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["beat-unit"] = make([]*MXML, len(r.BeatUnit))
	for i, c := range r.BeatUnit {
		children["beat-unit"][i] = c.ToMXML()
	}

	children["beat-unit-dot"] = make([]*MXML, len(r.BeatUnitDot))
	for i, c := range r.BeatUnitDot {
		children["beat-unit-dot"][i] = c.ToMXML()
	}

	return newMXML("beat-unit-tied", strings.TrimSpace(r.IValue), attributes, children)
}
