package musicxml

import (
	"encoding/xml"
	"strings"
)

// ScoreTimewise is the structure for score-timewise MusicXML element.
type ScoreTimewise struct {
	XMLName xml.Name `xml:"score-timewise"`

	VersionAttr string `xml:"version,attr,omitempty"`

	Credit         []Credit         `xml:"credit,omitempty"`
	Defaults       []Defaults       `xml:"defaults,omitempty"`
	Identification []Identification `xml:"identification,omitempty"`
	Measure        []Measure        `xml:"measure,omitempty"`
	MovementNumber []MovementNumber `xml:"movement-number,omitempty"`
	MovementTitle  []MovementTitle  `xml:"movement-title,omitempty"`
	PartList       []PartList       `xml:"part-list,omitempty"`
	Work           []Work           `xml:"work,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *ScoreTimewise) ToMXML() *MXML {
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

	children["measure"] = make([]*MXML, len(r.Measure))
	for i, c := range r.Measure {
		children["measure"][i] = c.ToMXML()
	}

	children["movement-number"] = make([]*MXML, len(r.MovementNumber))
	for i, c := range r.MovementNumber {
		children["movement-number"][i] = c.ToMXML()
	}

	children["movement-title"] = make([]*MXML, len(r.MovementTitle))
	for i, c := range r.MovementTitle {
		children["movement-title"][i] = c.ToMXML()
	}

	children["part-list"] = make([]*MXML, len(r.PartList))
	for i, c := range r.PartList {
		children["part-list"][i] = c.ToMXML()
	}

	children["work"] = make([]*MXML, len(r.Work))
	for i, c := range r.Work {
		children["work"][i] = c.ToMXML()
	}

	return newMXML("score-timewise", strings.TrimSpace(r.IValue), attributes, children)
}
