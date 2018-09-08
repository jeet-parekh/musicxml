package mxml

import (
	"encoding/xml"
	"strings"
)

// SystemLayout is the structure for system-layout MusicXML element.
type SystemLayout struct {
	XMLName xml.Name `xml:"system-layout"`

	SystemDistance    []SystemDistance    `xml:"system-distance,omitempty"`
	SystemDividers    []SystemDividers    `xml:"system-dividers,omitempty"`
	SystemMargins     []SystemMargins     `xml:"system-margins,omitempty"`
	TopSystemDistance []TopSystemDistance `xml:"top-system-distance,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *SystemLayout) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["system-distance"] = make([]*MXML, len(r.SystemDistance))
	for i, c := range r.SystemDistance {
		children["system-distance"][i] = c.ToMXML()
	}

	children["system-dividers"] = make([]*MXML, len(r.SystemDividers))
	for i, c := range r.SystemDividers {
		children["system-dividers"][i] = c.ToMXML()
	}

	children["system-margins"] = make([]*MXML, len(r.SystemMargins))
	for i, c := range r.SystemMargins {
		children["system-margins"][i] = c.ToMXML()
	}

	children["top-system-distance"] = make([]*MXML, len(r.TopSystemDistance))
	for i, c := range r.TopSystemDistance {
		children["top-system-distance"][i] = c.ToMXML()
	}

	return newMXML("system-layout", strings.TrimSpace(r.IValue), attributes, children)
}
