package musicxml

import (
	"encoding/xml"
	"strings"
)

// Rest is the structure for rest MusicXML element.
type Rest struct {
	XMLName xml.Name `xml:"rest"`

	MeasureAttr string `xml:"measure,attr,omitempty"`

	DisplayOctave []DisplayOctave `xml:"display-octave,omitempty"`
	DisplayStep   []DisplayStep   `xml:"display-step,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Rest) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["measure"] = r.MeasureAttr

	children["display-octave"] = make([]*MXML, len(r.DisplayOctave))
	for i, c := range r.DisplayOctave {
		children["display-octave"][i] = c.ToMXML()
	}

	children["display-step"] = make([]*MXML, len(r.DisplayStep))
	for i, c := range r.DisplayStep {
		children["display-step"][i] = c.ToMXML()
	}

	return newMXML("rest", strings.TrimSpace(r.IValue), attributes, children)
}
