package musicxml

import (
	"encoding/xml"
	"strings"
)

// Staff is the structure for staff MusicXML element.
type Staff struct {
	XMLName xml.Name `xml:"staff"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Staff) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("staff", strings.TrimSpace(r.IValue), attributes, children)
}
