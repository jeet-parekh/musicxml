package musicxml

import (
	"encoding/xml"
	"strings"
)

// Root is the structure for root MusicXML element.
type Root struct {
	XMLName xml.Name `xml:"root"`

	RootAlter []RootAlter `xml:"root-alter,omitempty"`
	RootStep  []RootStep  `xml:"root-step,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Root) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["root-alter"] = make([]*MXML, len(r.RootAlter))
	for i, c := range r.RootAlter {
		children["root-alter"][i] = c.ToMXML()
	}

	children["root-step"] = make([]*MXML, len(r.RootStep))
	for i, c := range r.RootStep {
		children["root-step"][i] = c.ToMXML()
	}

	return newMXML("root", strings.TrimSpace(r.IValue), attributes, children)
}
