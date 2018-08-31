package musicxml

import (
	"encoding/xml"
	"strings"
)

// StaffLines is the structure for staff-lines MusicXML element.
type StaffLines struct {
	XMLName xml.Name `xml:"staff-lines"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *StaffLines) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("staff-lines", strings.TrimSpace(r.IValue), attributes, children)
}
