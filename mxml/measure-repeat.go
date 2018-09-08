package mxml

import (
	"encoding/xml"
	"strings"
)

// MeasureRepeat is the structure for measure-repeat MusicXML element.
type MeasureRepeat struct {
	XMLName xml.Name `xml:"measure-repeat"`

	SlashesAttr string `xml:"slashes,attr,omitempty"`
	TypeAttr    string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MeasureRepeat) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["slashes"] = r.SlashesAttr
	attributes["type"] = r.TypeAttr

	return newMXML("measure-repeat", strings.TrimSpace(r.IValue), attributes, children)
}
