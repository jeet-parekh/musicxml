package mxml

import (
	"encoding/xml"
	"strings"
)

// Time is the structure for time MusicXML element.
type Time struct {
	XMLName xml.Name `xml:"time"`

	ColorAttr       string `xml:"color,attr,omitempty"`
	DefaultXAttr    string `xml:"default-x,attr,omitempty"`
	DefaultYAttr    string `xml:"default-y,attr,omitempty"`
	FontFamilyAttr  string `xml:"font-family,attr,omitempty"`
	FontSizeAttr    string `xml:"font-size,attr,omitempty"`
	FontStyleAttr   string `xml:"font-style,attr,omitempty"`
	FontWeightAttr  string `xml:"font-weight,attr,omitempty"`
	HalignAttr      string `xml:"halign,attr,omitempty"`
	IdAttr          string `xml:"id,attr,omitempty"`
	NumberAttr      string `xml:"number,attr,omitempty"`
	PrintObjectAttr string `xml:"print-object,attr,omitempty"`
	RelativeXAttr   string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr   string `xml:"relative-y,attr,omitempty"`
	SeparatorAttr   string `xml:"separator,attr,omitempty"`
	SymbolAttr      string `xml:"symbol,attr,omitempty"`
	ValignAttr      string `xml:"valign,attr,omitempty"`

	BeatType        []BeatType        `xml:"beat-type,omitempty"`
	Beats           []Beats           `xml:"beats,omitempty"`
	Interchangeable []Interchangeable `xml:"interchangeable,omitempty"`
	SenzaMisura     []SenzaMisura     `xml:"senza-misura,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Time) ToMXML() *MXML {
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
	attributes["number"] = r.NumberAttr
	attributes["print-object"] = r.PrintObjectAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["separator"] = r.SeparatorAttr
	attributes["symbol"] = r.SymbolAttr
	attributes["valign"] = r.ValignAttr

	children["beat-type"] = make([]*MXML, len(r.BeatType))
	for i, c := range r.BeatType {
		children["beat-type"][i] = c.ToMXML()
	}

	children["beats"] = make([]*MXML, len(r.Beats))
	for i, c := range r.Beats {
		children["beats"][i] = c.ToMXML()
	}

	children["interchangeable"] = make([]*MXML, len(r.Interchangeable))
	for i, c := range r.Interchangeable {
		children["interchangeable"][i] = c.ToMXML()
	}

	children["senza-misura"] = make([]*MXML, len(r.SenzaMisura))
	for i, c := range r.SenzaMisura {
		children["senza-misura"][i] = c.ToMXML()
	}

	return newMXML("time", strings.TrimSpace(r.IValue), attributes, children)
}
