package mxml

import (
	"encoding/xml"
	"strings"
)

// StaffDetails is the structure for staff-details MusicXML element.
type StaffDetails struct {
	XMLName xml.Name `xml:"staff-details"`

	NumberAttr       string `xml:"number,attr,omitempty"`
	PrintObjectAttr  string `xml:"print-object,attr,omitempty"`
	PrintSpacingAttr string `xml:"print-spacing,attr,omitempty"`
	ShowFretsAttr    string `xml:"show-frets,attr,omitempty"`

	Capo        []Capo        `xml:"capo,omitempty"`
	StaffLines  []StaffLines  `xml:"staff-lines,omitempty"`
	StaffSize   []StaffSize   `xml:"staff-size,omitempty"`
	StaffTuning []StaffTuning `xml:"staff-tuning,omitempty"`
	StaffType   []StaffType   `xml:"staff-type,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *StaffDetails) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["number"] = r.NumberAttr
	attributes["print-object"] = r.PrintObjectAttr
	attributes["print-spacing"] = r.PrintSpacingAttr
	attributes["show-frets"] = r.ShowFretsAttr

	children["capo"] = make([]*MXML, len(r.Capo))
	for i, c := range r.Capo {
		children["capo"][i] = c.ToMXML()
	}

	children["staff-lines"] = make([]*MXML, len(r.StaffLines))
	for i, c := range r.StaffLines {
		children["staff-lines"][i] = c.ToMXML()
	}

	children["staff-size"] = make([]*MXML, len(r.StaffSize))
	for i, c := range r.StaffSize {
		children["staff-size"][i] = c.ToMXML()
	}

	children["staff-tuning"] = make([]*MXML, len(r.StaffTuning))
	for i, c := range r.StaffTuning {
		children["staff-tuning"][i] = c.ToMXML()
	}

	children["staff-type"] = make([]*MXML, len(r.StaffType))
	for i, c := range r.StaffType {
		children["staff-type"][i] = c.ToMXML()
	}

	return newMXML("staff-details", strings.TrimSpace(r.IValue), attributes, children)
}
