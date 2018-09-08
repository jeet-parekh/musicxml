package mxml

import (
	"encoding/xml"
	"strings"
)

// TimeRelation is the structure for time-relation MusicXML element.
type TimeRelation struct {
	XMLName xml.Name `xml:"time-relation"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *TimeRelation) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("time-relation", strings.TrimSpace(r.IValue), attributes, children)
}
