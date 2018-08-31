package musicxml

import (
	"encoding/xml"
	"strings"
)

// OtherDynamics is the structure for other-dynamics MusicXML element.
type OtherDynamics struct {
	XMLName xml.Name `xml:"other-dynamics"`

	SmuflAttr string `xml:"smufl,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *OtherDynamics) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["smufl"] = r.SmuflAttr

	return newMXML("other-dynamics", strings.TrimSpace(r.IValue), attributes, children)
}
