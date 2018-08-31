package musicxml

import (
	"encoding/xml"
	"strings"
)

// Score is the structure for score MusicXML element.
type Score struct {
	XMLName xml.Name `xml:"score"`

	ActuateAttr string `xml:"actuate,attr,omitempty"`
	HrefAttr    string `xml:"href,attr,omitempty"`
	NewPageAttr string `xml:"new-page,attr,omitempty"`
	RoleAttr    string `xml:"role,attr,omitempty"`
	ShowAttr    string `xml:"show,attr,omitempty"`
	TitleAttr   string `xml:"title,attr,omitempty"`
	TypeAttr    string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Score) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["actuate"] = r.ActuateAttr
	attributes["href"] = r.HrefAttr
	attributes["new-page"] = r.NewPageAttr
	attributes["role"] = r.RoleAttr
	attributes["show"] = r.ShowAttr
	attributes["title"] = r.TitleAttr
	attributes["type"] = r.TypeAttr

	return newMXML("score", strings.TrimSpace(r.IValue), attributes, children)
}
