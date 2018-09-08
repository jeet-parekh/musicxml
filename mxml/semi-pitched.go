package mxml

import (
	"encoding/xml"
	"strings"
)

// SemiPitched is the structure for semi-pitched MusicXML element.
type SemiPitched struct {
	XMLName xml.Name `xml:"semi-pitched"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *SemiPitched) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("semi-pitched", strings.TrimSpace(r.IValue), attributes, children)
}
