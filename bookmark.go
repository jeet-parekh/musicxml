package musicxml

import (
	"encoding/xml"
	"strings"
)

// Bookmark is the structure for bookmark MusicXML element.
type Bookmark struct {
	XMLName xml.Name `xml:"bookmark"`

	ElementAttr  string `xml:"element,attr,omitempty"`
	IdAttr       string `xml:"id,attr,omitempty"`
	NameAttr     string `xml:"name,attr,omitempty"`
	PositionAttr string `xml:"position,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Bookmark) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["element"] = r.ElementAttr
	attributes["id"] = r.IdAttr
	attributes["name"] = r.NameAttr
	attributes["position"] = r.PositionAttr

	return newMXML("bookmark", strings.TrimSpace(r.IValue), attributes, children)
}
