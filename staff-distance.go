package musicxml

import (
	"encoding/xml"
	"strings"
)

// StaffDistance is the structure for staff-distance MusicXML element.
type StaffDistance struct {
	XMLName xml.Name `xml:"staff-distance"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *StaffDistance) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("staff-distance", strings.TrimSpace(r.IValue), attributes, children)
}
