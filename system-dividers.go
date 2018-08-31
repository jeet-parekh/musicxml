package musicxml

import (
	"encoding/xml"
	"strings"
)

// SystemDividers is the structure for system-dividers MusicXML element.
type SystemDividers struct {
	XMLName xml.Name `xml:"system-dividers"`

	LeftDivider  []LeftDivider  `xml:"left-divider,omitempty"`
	RightDivider []RightDivider `xml:"right-divider,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *SystemDividers) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["left-divider"] = make([]*MXML, len(r.LeftDivider))
	for i, c := range r.LeftDivider {
		children["left-divider"][i] = c.ToMXML()
	}

	children["right-divider"] = make([]*MXML, len(r.RightDivider))
	for i, c := range r.RightDivider {
		children["right-divider"][i] = c.ToMXML()
	}

	return newMXML("system-dividers", strings.TrimSpace(r.IValue), attributes, children)
}
