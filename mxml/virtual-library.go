package mxml

import (
	"encoding/xml"
	"strings"
)

// VirtualLibrary is the structure for virtual-library MusicXML element.
type VirtualLibrary struct {
	XMLName xml.Name `xml:"virtual-library"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *VirtualLibrary) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("virtual-library", strings.TrimSpace(r.IValue), attributes, children)
}
