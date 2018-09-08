package mxml

import (
	"encoding/xml"
	"strings"
)

// NormalType is the structure for normal-type MusicXML element.
type NormalType struct {
	XMLName xml.Name `xml:"normal-type"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *NormalType) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("normal-type", strings.TrimSpace(r.IValue), attributes, children)
}
