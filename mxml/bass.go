package mxml

import (
	"encoding/xml"
	"strings"
)

// Bass is the structure for bass MusicXML element.
type Bass struct {
	XMLName xml.Name `xml:"bass"`

	BassAlter []BassAlter `xml:"bass-alter,omitempty"`
	BassStep  []BassStep  `xml:"bass-step,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Bass) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["bass-alter"] = make([]*MXML, len(r.BassAlter))
	for i, c := range r.BassAlter {
		children["bass-alter"][i] = c.ToMXML()
	}

	children["bass-step"] = make([]*MXML, len(r.BassStep))
	for i, c := range r.BassStep {
		children["bass-step"][i] = c.ToMXML()
	}

	return newMXML("bass", strings.TrimSpace(r.IValue), attributes, children)
}
