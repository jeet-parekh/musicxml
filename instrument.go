package musicxml

import (
	"encoding/xml"
	"strings"
)

// Instrument is the structure for instrument MusicXML element.
type Instrument struct {
	XMLName xml.Name `xml:"instrument"`

	IdAttr string `xml:"id,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Instrument) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["id"] = r.IdAttr

	return newMXML("instrument", strings.TrimSpace(r.IValue), attributes, children)
}
