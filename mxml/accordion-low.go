package mxml

import (
	"encoding/xml"
	"strings"
)

// AccordionLow is the structure for accordion-low MusicXML element.
type AccordionLow struct {
	XMLName xml.Name `xml:"accordion-low"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *AccordionLow) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("accordion-low", strings.TrimSpace(r.IValue), attributes, children)
}
