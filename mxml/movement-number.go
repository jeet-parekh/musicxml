package mxml

import (
	"encoding/xml"
	"strings"
)

// MovementNumber is the structure for movement-number MusicXML element.
type MovementNumber struct {
	XMLName xml.Name `xml:"movement-number"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MovementNumber) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("movement-number", strings.TrimSpace(r.IValue), attributes, children)
}
