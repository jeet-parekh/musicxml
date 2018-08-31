package musicxml

import (
	"encoding/xml"
	"strings"
)

// StaffSize is the structure for staff-size MusicXML element.
type StaffSize struct {
	XMLName xml.Name `xml:"staff-size"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *StaffSize) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("staff-size", strings.TrimSpace(r.IValue), attributes, children)
}
