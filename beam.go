package musicxml

import (
	"encoding/xml"
	"strings"
)

// Beam is the structure for beam MusicXML element.
type Beam struct {
	XMLName xml.Name `xml:"beam"`

	ColorAttr    string `xml:"color,attr,omitempty"`
	FanAttr      string `xml:"fan,attr,omitempty"`
	IdAttr       string `xml:"id,attr,omitempty"`
	NumberAttr   string `xml:"number,attr,omitempty"`
	RepeaterAttr string `xml:"repeater,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Beam) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["fan"] = r.FanAttr
	attributes["id"] = r.IdAttr
	attributes["number"] = r.NumberAttr
	attributes["repeater"] = r.RepeaterAttr

	return newMXML("beam", strings.TrimSpace(r.IValue), attributes, children)
}
