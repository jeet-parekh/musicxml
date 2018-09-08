package mxml

import (
	"encoding/xml"
	"strings"
)

// OtherPercussion is the structure for other-percussion MusicXML element.
type OtherPercussion struct {
	XMLName xml.Name `xml:"other-percussion"`

	SmuflAttr string `xml:"smufl,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *OtherPercussion) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["smufl"] = r.SmuflAttr

	return newMXML("other-percussion", strings.TrimSpace(r.IValue), attributes, children)
}
