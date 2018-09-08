package mxml

import (
	"encoding/xml"
	"strings"
)

// PartList is the structure for part-list MusicXML element.
type PartList struct {
	XMLName xml.Name `xml:"part-list"`

	PartGroup []PartGroup `xml:"part-group,omitempty"`
	ScorePart []ScorePart `xml:"score-part,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *PartList) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["part-group"] = make([]*MXML, len(r.PartGroup))
	for i, c := range r.PartGroup {
		children["part-group"][i] = c.ToMXML()
	}

	children["score-part"] = make([]*MXML, len(r.ScorePart))
	for i, c := range r.ScorePart {
		children["score-part"][i] = c.ToMXML()
	}

	return newMXML("part-list", strings.TrimSpace(r.IValue), attributes, children)
}
