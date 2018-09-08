package mxml

import (
	"encoding/xml"
	"strings"
)

// Ensemble is the structure for ensemble MusicXML element.
type Ensemble struct {
	XMLName xml.Name `xml:"ensemble"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Ensemble) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("ensemble", strings.TrimSpace(r.IValue), attributes, children)
}
