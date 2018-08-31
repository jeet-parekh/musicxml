package musicxml

import (
	"encoding/xml"
	"strings"
)

// MidiDevice is the structure for midi-device MusicXML element.
type MidiDevice struct {
	XMLName xml.Name `xml:"midi-device"`

	IdAttr   string `xml:"id,attr,omitempty"`
	PortAttr string `xml:"port,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MidiDevice) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["id"] = r.IdAttr
	attributes["port"] = r.PortAttr

	return newMXML("midi-device", strings.TrimSpace(r.IValue), attributes, children)
}
