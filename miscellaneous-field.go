package musicxml

import (
	"encoding/xml"
	"strings"
)

// MiscellaneousField is the structure for miscellaneous-field MusicXML element.
type MiscellaneousField struct {
	XMLName xml.Name `xml:"miscellaneous-field"`

	NameAttr string `xml:"name,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MiscellaneousField) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["name"] = r.NameAttr

	return newMXML("miscellaneous-field", strings.TrimSpace(r.IValue), attributes, children)
}
