package musicxml

import (
	"encoding/xml"
	"strings"
)

// PartGroup is the structure for part-group MusicXML element.
type PartGroup struct {
	XMLName xml.Name `xml:"part-group"`

	NumberAttr string `xml:"number,attr,omitempty"`
	TypeAttr   string `xml:"type,attr,omitempty"`

	Footnote                 []Footnote                 `xml:"footnote,omitempty"`
	GroupAbbreviation        []GroupAbbreviation        `xml:"group-abbreviation,omitempty"`
	GroupAbbreviationDisplay []GroupAbbreviationDisplay `xml:"group-abbreviation-display,omitempty"`
	GroupBarline             []GroupBarline             `xml:"group-barline,omitempty"`
	GroupName                []GroupName                `xml:"group-name,omitempty"`
	GroupNameDisplay         []GroupNameDisplay         `xml:"group-name-display,omitempty"`
	GroupSymbol              []GroupSymbol              `xml:"group-symbol,omitempty"`
	GroupTime                []GroupTime                `xml:"group-time,omitempty"`
	Level                    []Level                    `xml:"level,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *PartGroup) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["number"] = r.NumberAttr
	attributes["type"] = r.TypeAttr

	children["footnote"] = make([]*MXML, len(r.Footnote))
	for i, c := range r.Footnote {
		children["footnote"][i] = c.ToMXML()
	}

	children["group-abbreviation"] = make([]*MXML, len(r.GroupAbbreviation))
	for i, c := range r.GroupAbbreviation {
		children["group-abbreviation"][i] = c.ToMXML()
	}

	children["group-abbreviation-display"] = make([]*MXML, len(r.GroupAbbreviationDisplay))
	for i, c := range r.GroupAbbreviationDisplay {
		children["group-abbreviation-display"][i] = c.ToMXML()
	}

	children["group-barline"] = make([]*MXML, len(r.GroupBarline))
	for i, c := range r.GroupBarline {
		children["group-barline"][i] = c.ToMXML()
	}

	children["group-name"] = make([]*MXML, len(r.GroupName))
	for i, c := range r.GroupName {
		children["group-name"][i] = c.ToMXML()
	}

	children["group-name-display"] = make([]*MXML, len(r.GroupNameDisplay))
	for i, c := range r.GroupNameDisplay {
		children["group-name-display"][i] = c.ToMXML()
	}

	children["group-symbol"] = make([]*MXML, len(r.GroupSymbol))
	for i, c := range r.GroupSymbol {
		children["group-symbol"][i] = c.ToMXML()
	}

	children["group-time"] = make([]*MXML, len(r.GroupTime))
	for i, c := range r.GroupTime {
		children["group-time"][i] = c.ToMXML()
	}

	children["level"] = make([]*MXML, len(r.Level))
	for i, c := range r.Level {
		children["level"][i] = c.ToMXML()
	}

	return newMXML("part-group", strings.TrimSpace(r.IValue), attributes, children)
}
