package mxml

import (
	"encoding/xml"
	"strings"
)

// Print is the structure for print MusicXML element.
type Print struct {
	XMLName xml.Name `xml:"print"`

	BlankPageAttr    string `xml:"blank-page,attr,omitempty"`
	IdAttr           string `xml:"id,attr,omitempty"`
	NewPageAttr      string `xml:"new-page,attr,omitempty"`
	NewSystemAttr    string `xml:"new-system,attr,omitempty"`
	PageNumberAttr   string `xml:"page-number,attr,omitempty"`
	StaffSpacingAttr string `xml:"staff-spacing,attr,omitempty"`

	MeasureLayout           []MeasureLayout           `xml:"measure-layout,omitempty"`
	MeasureNumbering        []MeasureNumbering        `xml:"measure-numbering,omitempty"`
	PageLayout              []PageLayout              `xml:"page-layout,omitempty"`
	PartAbbreviationDisplay []PartAbbreviationDisplay `xml:"part-abbreviation-display,omitempty"`
	PartNameDisplay         []PartNameDisplay         `xml:"part-name-display,omitempty"`
	StaffLayout             []StaffLayout             `xml:"staff-layout,omitempty"`
	SystemLayout            []SystemLayout            `xml:"system-layout,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Print) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["blank-page"] = r.BlankPageAttr
	attributes["id"] = r.IdAttr
	attributes["new-page"] = r.NewPageAttr
	attributes["new-system"] = r.NewSystemAttr
	attributes["page-number"] = r.PageNumberAttr
	attributes["staff-spacing"] = r.StaffSpacingAttr

	children["measure-layout"] = make([]*MXML, len(r.MeasureLayout))
	for i, c := range r.MeasureLayout {
		children["measure-layout"][i] = c.ToMXML()
	}

	children["measure-numbering"] = make([]*MXML, len(r.MeasureNumbering))
	for i, c := range r.MeasureNumbering {
		children["measure-numbering"][i] = c.ToMXML()
	}

	children["page-layout"] = make([]*MXML, len(r.PageLayout))
	for i, c := range r.PageLayout {
		children["page-layout"][i] = c.ToMXML()
	}

	children["part-abbreviation-display"] = make([]*MXML, len(r.PartAbbreviationDisplay))
	for i, c := range r.PartAbbreviationDisplay {
		children["part-abbreviation-display"][i] = c.ToMXML()
	}

	children["part-name-display"] = make([]*MXML, len(r.PartNameDisplay))
	for i, c := range r.PartNameDisplay {
		children["part-name-display"][i] = c.ToMXML()
	}

	children["staff-layout"] = make([]*MXML, len(r.StaffLayout))
	for i, c := range r.StaffLayout {
		children["staff-layout"][i] = c.ToMXML()
	}

	children["system-layout"] = make([]*MXML, len(r.SystemLayout))
	for i, c := range r.SystemLayout {
		children["system-layout"][i] = c.ToMXML()
	}

	return newMXML("print", strings.TrimSpace(r.IValue), attributes, children)
}
