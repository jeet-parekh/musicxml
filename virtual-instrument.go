package musicxml

import (
	"encoding/xml"
	"strings"
)

// VirtualInstrument is the structure for virtual-instrument MusicXML element.
type VirtualInstrument struct {
	XMLName xml.Name `xml:"virtual-instrument"`

	VirtualLibrary []VirtualLibrary `xml:"virtual-library,omitempty"`
	VirtualName    []VirtualName    `xml:"virtual-name,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *VirtualInstrument) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["virtual-library"] = make([]*MXML, len(r.VirtualLibrary))
	for i, c := range r.VirtualLibrary {
		children["virtual-library"][i] = c.ToMXML()
	}

	children["virtual-name"] = make([]*MXML, len(r.VirtualName))
	for i, c := range r.VirtualName {
		children["virtual-name"][i] = c.ToMXML()
	}

	return newMXML("virtual-instrument", strings.TrimSpace(r.IValue), attributes, children)
}
