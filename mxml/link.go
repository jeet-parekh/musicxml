package mxml

import (
	"encoding/xml"
	"strings"
)

// Link is the structure for link MusicXML element.
type Link struct {
	XMLName xml.Name `xml:"link"`

	ActuateAttr   string `xml:"actuate,attr,omitempty"`
	DefaultXAttr  string `xml:"default-x,attr,omitempty"`
	DefaultYAttr  string `xml:"default-y,attr,omitempty"`
	ElementAttr   string `xml:"element,attr,omitempty"`
	HrefAttr      string `xml:"href,attr,omitempty"`
	NameAttr      string `xml:"name,attr,omitempty"`
	PositionAttr  string `xml:"position,attr,omitempty"`
	RelativeXAttr string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr string `xml:"relative-y,attr,omitempty"`
	RoleAttr      string `xml:"role,attr,omitempty"`
	ShowAttr      string `xml:"show,attr,omitempty"`
	TitleAttr     string `xml:"title,attr,omitempty"`
	TypeAttr      string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Link) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["actuate"] = r.ActuateAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["element"] = r.ElementAttr
	attributes["href"] = r.HrefAttr
	attributes["name"] = r.NameAttr
	attributes["position"] = r.PositionAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["role"] = r.RoleAttr
	attributes["show"] = r.ShowAttr
	attributes["title"] = r.TitleAttr
	attributes["type"] = r.TypeAttr

	return newMXML("link", strings.TrimSpace(r.IValue), attributes, children)
}
