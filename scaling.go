package musicxml

import (
	"encoding/xml"
	"strings"
)

// Scaling is the structure for scaling MusicXML element.
type Scaling struct {
	XMLName xml.Name `xml:"scaling"`

	Millimeters []Millimeters `xml:"millimeters,omitempty"`
	Tenths      []Tenths      `xml:"tenths,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Scaling) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["millimeters"] = make([]*MXML, len(r.Millimeters))
	for i, c := range r.Millimeters {
		children["millimeters"][i] = c.ToMXML()
	}

	children["tenths"] = make([]*MXML, len(r.Tenths))
	for i, c := range r.Tenths {
		children["tenths"][i] = c.ToMXML()
	}

	return newMXML("scaling", strings.TrimSpace(r.IValue), attributes, children)
}
