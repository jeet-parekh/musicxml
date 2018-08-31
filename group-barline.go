package musicxml

import (
	"encoding/xml"
	"strings"
)

// GroupBarline is the structure for group-barline MusicXML element.
type GroupBarline struct {
	XMLName xml.Name `xml:"group-barline"`

	ColorAttr string `xml:"color,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *GroupBarline) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr

	return newMXML("group-barline", strings.TrimSpace(r.IValue), attributes, children)
}
