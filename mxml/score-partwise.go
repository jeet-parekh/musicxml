package mxml

import (
	"encoding/xml"
	"strings"
)

// ScorePartwise is the structure for score-partwise MusicXML element.
type ScorePartwise struct {
	XMLName xml.Name `xml:"score-partwise"`

	VersionAttr string `xml:"version,attr,omitempty"`

	Credit         []Credit         `xml:"credit,omitempty"`
	Defaults       []Defaults       `xml:"defaults,omitempty"`
	Identification []Identification `xml:"identification,omitempty"`
	MovementNumber []MovementNumber `xml:"movement-number,omitempty"`
	MovementTitle  []MovementTitle  `xml:"movement-title,omitempty"`
	Part           []Part           `xml:"part,omitempty"`
	PartList       []PartList       `xml:"part-list,omitempty"`
	Work           []Work           `xml:"work,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *ScorePartwise) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["version"] = r.VersionAttr

	children["credit"] = make([]*MXML, len(r.Credit))
	for i, c := range r.Credit {
		children["credit"][i] = c.ToMXML()
	}

	children["defaults"] = make([]*MXML, len(r.Defaults))
	for i, c := range r.Defaults {
		children["defaults"][i] = c.ToMXML()
	}

	children["identification"] = make([]*MXML, len(r.Identification))
	for i, c := range r.Identification {
		children["identification"][i] = c.ToMXML()
	}

	children["movement-number"] = make([]*MXML, len(r.MovementNumber))
	for i, c := range r.MovementNumber {
		children["movement-number"][i] = c.ToMXML()
	}

	children["movement-title"] = make([]*MXML, len(r.MovementTitle))
	for i, c := range r.MovementTitle {
		children["movement-title"][i] = c.ToMXML()
	}

	children["part"] = make([]*MXML, len(r.Part))
	for i, c := range r.Part {
		children["part"][i] = c.ToMXML()
	}

	children["part-list"] = make([]*MXML, len(r.PartList))
	for i, c := range r.PartList {
		children["part-list"][i] = c.ToMXML()
	}

	children["work"] = make([]*MXML, len(r.Work))
	for i, c := range r.Work {
		children["work"][i] = c.ToMXML()
	}

	return newMXML("score-partwise", strings.TrimSpace(r.IValue), attributes, children)
}
