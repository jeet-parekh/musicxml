package mxml

import (
	"encoding/xml"
	"strings"
)

// Glass is the structure for glass MusicXML element.
type Glass struct {
	XMLName xml.Name `xml:"glass"`

	SmuflAttr string `xml:"smufl,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Glass) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["smufl"] = r.SmuflAttr

	return newMXML("glass", strings.TrimSpace(r.IValue), attributes, children)
}
