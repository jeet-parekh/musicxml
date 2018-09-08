package mxml

import (
	"encoding/xml"
	"strings"
)

// MovementTitle is the structure for movement-title MusicXML element.
type MovementTitle struct {
	XMLName xml.Name `xml:"movement-title"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MovementTitle) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("movement-title", strings.TrimSpace(r.IValue), attributes, children)
}
