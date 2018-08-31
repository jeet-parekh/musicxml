package musicxml

import (
	"encoding/xml"
	"strings"
)

// StaffLayout is the structure for staff-layout MusicXML element.
type StaffLayout struct {
	XMLName xml.Name `xml:"staff-layout"`

	NumberAttr string `xml:"number,attr,omitempty"`

	StaffDistance []StaffDistance `xml:"staff-distance,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *StaffLayout) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["number"] = r.NumberAttr

	children["staff-distance"] = make([]*MXML, len(r.StaffDistance))
	for i, c := range r.StaffDistance {
		children["staff-distance"][i] = c.ToMXML()
	}

	return newMXML("staff-layout", strings.TrimSpace(r.IValue), attributes, children)
}
