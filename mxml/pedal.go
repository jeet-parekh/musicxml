package mxml

import (
	"encoding/xml"
	"strings"
)

// Pedal is the structure for pedal MusicXML element.
type Pedal struct {
	XMLName xml.Name `xml:"pedal"`

	AbbreviatedAttr string `xml:"abbreviated,attr,omitempty"`
	ColorAttr       string `xml:"color,attr,omitempty"`
	DefaultXAttr    string `xml:"default-x,attr,omitempty"`
	DefaultYAttr    string `xml:"default-y,attr,omitempty"`
	FontFamilyAttr  string `xml:"font-family,attr,omitempty"`
	FontSizeAttr    string `xml:"font-size,attr,omitempty"`
	FontStyleAttr   string `xml:"font-style,attr,omitempty"`
	FontWeightAttr  string `xml:"font-weight,attr,omitempty"`
	HalignAttr      string `xml:"halign,attr,omitempty"`
	IdAttr          string `xml:"id,attr,omitempty"`
	LineAttr        string `xml:"line,attr,omitempty"`
	NumberAttr      string `xml:"number,attr,omitempty"`
	RelativeXAttr   string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr   string `xml:"relative-y,attr,omitempty"`
	SignAttr        string `xml:"sign,attr,omitempty"`
	TypeAttr        string `xml:"type,attr,omitempty"`
	ValignAttr      string `xml:"valign,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Pedal) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["abbreviated"] = r.AbbreviatedAttr
	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["halign"] = r.HalignAttr
	attributes["id"] = r.IdAttr
	attributes["line"] = r.LineAttr
	attributes["number"] = r.NumberAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["sign"] = r.SignAttr
	attributes["type"] = r.TypeAttr
	attributes["valign"] = r.ValignAttr

	return newMXML("pedal", strings.TrimSpace(r.IValue), attributes, children)
}
