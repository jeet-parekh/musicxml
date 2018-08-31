package musicxml

import (
	"encoding/xml"
	"strings"
)

// VirtualName is the structure for virtual-name MusicXML element.
type VirtualName struct {
	XMLName xml.Name `xml:"virtual-name"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *VirtualName) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("virtual-name", strings.TrimSpace(r.IValue), attributes, children)
}
