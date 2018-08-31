package musicxml

import (
	"encoding/xml"
	"strings"
)

// SystemMargins is the structure for system-margins MusicXML element.
type SystemMargins struct {
	XMLName xml.Name `xml:"system-margins"`

	LeftMargin  []LeftMargin  `xml:"left-margin,omitempty"`
	RightMargin []RightMargin `xml:"right-margin,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *SystemMargins) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["left-margin"] = make([]*MXML, len(r.LeftMargin))
	for i, c := range r.LeftMargin {
		children["left-margin"][i] = c.ToMXML()
	}

	children["right-margin"] = make([]*MXML, len(r.RightMargin))
	for i, c := range r.RightMargin {
		children["right-margin"][i] = c.ToMXML()
	}

	return newMXML("system-margins", strings.TrimSpace(r.IValue), attributes, children)
}
