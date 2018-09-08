package mxml

import (
	"encoding/xml"
	"strings"
)

// GroupTime is the structure for group-time MusicXML element.
type GroupTime struct {
	XMLName xml.Name `xml:"group-time"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *GroupTime) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("group-time", strings.TrimSpace(r.IValue), attributes, children)
}
