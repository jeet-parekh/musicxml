package mxml

import (
	"encoding/xml"
	"strings"
)

// AccordionRegistration is the structure for accordion-registration MusicXML element.
type AccordionRegistration struct {
	XMLName xml.Name `xml:"accordion-registration"`

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
	ValignAttr     string `xml:"valign,attr,omitempty"`

	AccordionHigh   []AccordionHigh   `xml:"accordion-high,omitempty"`
	AccordionLow    []AccordionLow    `xml:"accordion-low,omitempty"`
	AccordionMiddle []AccordionMiddle `xml:"accordion-middle,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *AccordionRegistration) ToMXML() *MXML {
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
	attributes["valign"] = r.ValignAttr

	children["accordion-high"] = make([]*MXML, len(r.AccordionHigh))
	for i, c := range r.AccordionHigh {
		children["accordion-high"][i] = c.ToMXML()
	}

	children["accordion-low"] = make([]*MXML, len(r.AccordionLow))
	for i, c := range r.AccordionLow {
		children["accordion-low"][i] = c.ToMXML()
	}

	children["accordion-middle"] = make([]*MXML, len(r.AccordionMiddle))
	for i, c := range r.AccordionMiddle {
		children["accordion-middle"][i] = c.ToMXML()
	}

	return newMXML("accordion-registration", strings.TrimSpace(r.IValue), attributes, children)
}
