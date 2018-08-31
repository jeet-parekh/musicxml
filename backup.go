package musicxml

import (
	"encoding/xml"
	"strings"
)

// Backup is the structure for backup MusicXML element.
type Backup struct {
	XMLName xml.Name `xml:"backup"`

	Duration []Duration `xml:"duration,omitempty"`
	Footnote []Footnote `xml:"footnote,omitempty"`
	Level    []Level    `xml:"level,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Backup) ToMXML() *MXML {
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

	return newMXML("backup", strings.TrimSpace(r.IValue), attributes, children)
}
