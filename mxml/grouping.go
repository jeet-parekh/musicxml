package mxml

import (
	"encoding/xml"
	"strings"
)

// Grouping is the structure for grouping MusicXML element.
type Grouping struct {
	XMLName xml.Name `xml:"grouping"`

	IdAttr       string `xml:"id,attr,omitempty"`
	MemberOfAttr string `xml:"member-of,attr,omitempty"`
	NumberAttr   string `xml:"number,attr,omitempty"`
	TypeAttr     string `xml:"type,attr,omitempty"`

	Feature []Feature `xml:"feature,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Grouping) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["id"] = r.IdAttr
	attributes["member-of"] = r.MemberOfAttr
	attributes["number"] = r.NumberAttr
	attributes["type"] = r.TypeAttr

	children["feature"] = make([]*MXML, len(r.Feature))
	for i, c := range r.Feature {
		children["feature"][i] = c.ToMXML()
	}

	return newMXML("grouping", strings.TrimSpace(r.IValue), attributes, children)
}
