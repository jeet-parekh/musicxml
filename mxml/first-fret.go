package mxml

import (
	"encoding/xml"
	"strings"
)

// FirstFret is the structure for first-fret MusicXML element.
type FirstFret struct {
	XMLName xml.Name `xml:"first-fret"`

	LocationAttr string `xml:"location,attr,omitempty"`
	TextAttr     string `xml:"text,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *FirstFret) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["location"] = r.LocationAttr
	attributes["text"] = r.TextAttr

	return newMXML("first-fret", strings.TrimSpace(r.IValue), attributes, children)
}
