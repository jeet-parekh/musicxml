package mxml

import (
	"encoding/xml"
	"strings"
)

// Pitched is the structure for pitched MusicXML element.
type Pitched struct {
	XMLName xml.Name `xml:"pitched"`

	SmuflAttr string `xml:"smufl,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Pitched) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["smufl"] = r.SmuflAttr

	return newMXML("pitched", strings.TrimSpace(r.IValue), attributes, children)
}
