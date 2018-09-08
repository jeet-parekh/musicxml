package mxml

import (
	"encoding/xml"
	"strings"
)

// MeasureDistance is the structure for measure-distance MusicXML element.
type MeasureDistance struct {
	XMLName xml.Name `xml:"measure-distance"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MeasureDistance) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("measure-distance", strings.TrimSpace(r.IValue), attributes, children)
}
