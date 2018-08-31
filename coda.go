package musicxml

import (
	"encoding/xml"
	"strings"
)

// Coda is the structure for coda MusicXML element.
type Coda struct {
	XMLName xml.Name `xml:"coda"`

	ColorAttr      string `xml:"color,attr,omitempty"`
	DefaultXAttr   string `xml:"default-x,attr,omitempty"`
	DefaultYAttr   string `xml:"default-y,attr,omitempty"`
	FontFamilyAttr string `xml:"font-family,attr,omitempty"`
	FontSizeAttr   string `xml:"font-size,attr,omitempty"`
	FontStyleAttr  string `xml:"font-style,attr,omitempty"`
	FontWeightAttr string `xml:"font-weight,attr,omitempty"`
	HalignAttr     string `xml:"halign,attr,omitempty"`
	IdAttr         string `xml:"id,attr,omitempty"`
	RelativeXAttr  string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr  string `xml:"relative-y,attr,omitempty"`
	SmuflAttr      string `xml:"smufl,attr,omitempty"`
	ValignAttr     string `xml:"valign,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Coda) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["halign"] = r.HalignAttr
	attributes["id"] = r.IdAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["smufl"] = r.SmuflAttr
	attributes["valign"] = r.ValignAttr

	return newMXML("coda", strings.TrimSpace(r.IValue), attributes, children)
}
