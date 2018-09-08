package mxml

import (
	"encoding/xml"
	"strings"
)

// PageHeight is the structure for page-height MusicXML element.
type PageHeight struct {
	XMLName xml.Name `xml:"page-height"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *PageHeight) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("page-height", strings.TrimSpace(r.IValue), attributes, children)
}
