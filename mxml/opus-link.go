package mxml

import (
	"encoding/xml"
	"strings"
)

// OpusLink is the structure for opus-link MusicXML element.
type OpusLink struct {
	XMLName xml.Name `xml:"opus-link"`

	ActuateAttr string `xml:"actuate,attr,omitempty"`
	HrefAttr    string `xml:"href,attr,omitempty"`
	RoleAttr    string `xml:"role,attr,omitempty"`
	ShowAttr    string `xml:"show,attr,omitempty"`
	TitleAttr   string `xml:"title,attr,omitempty"`
	TypeAttr    string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *OpusLink) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["actuate"] = r.ActuateAttr
	attributes["href"] = r.HrefAttr
	attributes["role"] = r.RoleAttr
	attributes["show"] = r.ShowAttr
	attributes["title"] = r.TitleAttr
	attributes["type"] = r.TypeAttr

	return newMXML("opus-link", strings.TrimSpace(r.IValue), attributes, children)
}
