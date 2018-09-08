package mxml

import (
	"encoding/xml"
	"strings"
)

// AccordionMiddle is the structure for accordion-middle MusicXML element.
type AccordionMiddle struct {
	XMLName xml.Name `xml:"accordion-middle"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *AccordionMiddle) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("accordion-middle", strings.TrimSpace(r.IValue), attributes, children)
}
