package musicxml

import (
	"encoding/xml"
	"strings"
)

// Image is the structure for image MusicXML element.
type Image struct {
	XMLName xml.Name `xml:"image"`

	DefaultXAttr  string `xml:"default-x,attr,omitempty"`
	DefaultYAttr  string `xml:"default-y,attr,omitempty"`
	HalignAttr    string `xml:"halign,attr,omitempty"`
	HeightAttr    string `xml:"height,attr,omitempty"`
	IdAttr        string `xml:"id,attr,omitempty"`
	RelativeXAttr string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr string `xml:"relative-y,attr,omitempty"`
	SourceAttr    string `xml:"source,attr,omitempty"`
	TypeAttr      string `xml:"type,attr,omitempty"`
	ValignAttr    string `xml:"valign,attr,omitempty"`
	WidthAttr     string `xml:"width,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Image) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["halign"] = r.HalignAttr
	attributes["height"] = r.HeightAttr
	attributes["id"] = r.IdAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["source"] = r.SourceAttr
	attributes["type"] = r.TypeAttr
	attributes["valign"] = r.ValignAttr
	attributes["width"] = r.WidthAttr

	return newMXML("image", strings.TrimSpace(r.IValue), attributes, children)
}
