package musicxml

import (
	"encoding/xml"
	"strings"
)

// Unpitched is the structure for unpitched MusicXML element.
type Unpitched struct {
	XMLName xml.Name `xml:"unpitched"`

	DisplayOctave []DisplayOctave `xml:"display-octave,omitempty"`
	DisplayStep   []DisplayStep   `xml:"display-step,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Unpitched) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["display-octave"] = make([]*MXML, len(r.DisplayOctave))
	for i, c := range r.DisplayOctave {
		children["display-octave"][i] = c.ToMXML()
	}

	children["display-step"] = make([]*MXML, len(r.DisplayStep))
	for i, c := range r.DisplayStep {
		children["display-step"][i] = c.ToMXML()
	}

	return newMXML("unpitched", strings.TrimSpace(r.IValue), attributes, children)
}
