package musicxml

import (
	"encoding/xml"
	"strings"
)

// FrameNote is the structure for frame-note MusicXML element.
type FrameNote struct {
	XMLName xml.Name `xml:"frame-note"`

	Barre     []Barre     `xml:"barre,omitempty"`
	Fingering []Fingering `xml:"fingering,omitempty"`
	Fret      []Fret      `xml:"fret,omitempty"`
	String    []String    `xml:"string,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *FrameNote) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["barre"] = make([]*MXML, len(r.Barre))
	for i, c := range r.Barre {
		children["barre"][i] = c.ToMXML()
	}

	children["fingering"] = make([]*MXML, len(r.Fingering))
	for i, c := range r.Fingering {
		children["fingering"][i] = c.ToMXML()
	}

	children["fret"] = make([]*MXML, len(r.Fret))
	for i, c := range r.Fret {
		children["fret"][i] = c.ToMXML()
	}

	children["string"] = make([]*MXML, len(r.String))
	for i, c := range r.String {
		children["string"][i] = c.ToMXML()
	}

	return newMXML("frame-note", strings.TrimSpace(r.IValue), attributes, children)
}
