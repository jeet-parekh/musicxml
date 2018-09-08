package mxml

import (
	"encoding/xml"
	"strings"
)

// StaffType is the structure for staff-type MusicXML element.
type StaffType struct {
	XMLName xml.Name `xml:"staff-type"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *StaffType) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("staff-type", strings.TrimSpace(r.IValue), attributes, children)
}
