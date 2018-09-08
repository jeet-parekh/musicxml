package mxml

import (
	"encoding/xml"
	"strings"
)

// AccordionHigh is the structure for accordion-high MusicXML element.
type AccordionHigh struct {
	XMLName xml.Name `xml:"accordion-high"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *AccordionHigh) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("accordion-high", strings.TrimSpace(r.IValue), attributes, children)
}
