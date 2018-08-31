package musicxml

import (
	"encoding/xml"
	"strings"
)

// Forward is the structure for forward MusicXML element.
type Forward struct {
	XMLName xml.Name `xml:"forward"`

	Duration []Duration `xml:"duration,omitempty"`
	Footnote []Footnote `xml:"footnote,omitempty"`
	Level    []Level    `xml:"level,omitempty"`
	Staff    []Staff    `xml:"staff,omitempty"`
	Voice    []Voice    `xml:"voice,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Forward) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["duration"] = make([]*MXML, len(r.Duration))
	for i, c := range r.Duration {
		children["duration"][i] = c.ToMXML()
	}

	children["footnote"] = make([]*MXML, len(r.Footnote))
	for i, c := range r.Footnote {
		children["footnote"][i] = c.ToMXML()
	}

	children["level"] = make([]*MXML, len(r.Level))
	for i, c := range r.Level {
		children["level"][i] = c.ToMXML()
	}

	children["staff"] = make([]*MXML, len(r.Staff))
	for i, c := range r.Staff {
		children["staff"][i] = c.ToMXML()
	}

	children["voice"] = make([]*MXML, len(r.Voice))
	for i, c := range r.Voice {
		children["voice"][i] = c.ToMXML()
	}

	return newMXML("forward", strings.TrimSpace(r.IValue), attributes, children)
}
