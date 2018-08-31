package musicxml

import (
	"encoding/xml"
	"strings"
)

// BasePitch is the structure for base-pitch MusicXML element.
type BasePitch struct {
	XMLName xml.Name `xml:"base-pitch"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *BasePitch) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("base-pitch", strings.TrimSpace(r.IValue), attributes, children)
}
