package mxml

import (
	"encoding/xml"
	"strings"
)

// Miscellaneous is the structure for miscellaneous MusicXML element.
type Miscellaneous struct {
	XMLName xml.Name `xml:"miscellaneous"`

	MiscellaneousField []MiscellaneousField `xml:"miscellaneous-field,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Miscellaneous) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["miscellaneous-field"] = make([]*MXML, len(r.MiscellaneousField))
	for i, c := range r.MiscellaneousField {
		children["miscellaneous-field"][i] = c.ToMXML()
	}

	return newMXML("miscellaneous", strings.TrimSpace(r.IValue), attributes, children)
}
