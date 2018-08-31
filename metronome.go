package musicxml

import (
	"encoding/xml"
	"strings"
)

// Metronome is the structure for metronome MusicXML element.
type Metronome struct {
	XMLName xml.Name `xml:"metronome"`

	ColorAttr       string `xml:"color,attr,omitempty"`
	DefaultXAttr    string `xml:"default-x,attr,omitempty"`
	DefaultYAttr    string `xml:"default-y,attr,omitempty"`
	FontFamilyAttr  string `xml:"font-family,attr,omitempty"`
	FontSizeAttr    string `xml:"font-size,attr,omitempty"`
	FontStyleAttr   string `xml:"font-style,attr,omitempty"`
	FontWeightAttr  string `xml:"font-weight,attr,omitempty"`
	HalignAttr      string `xml:"halign,attr,omitempty"`
	IdAttr          string `xml:"id,attr,omitempty"`
	JustifyAttr     string `xml:"justify,attr,omitempty"`
	ParenthesesAttr string `xml:"parentheses,attr,omitempty"`
	RelativeXAttr   string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr   string `xml:"relative-y,attr,omitempty"`
	ValignAttr      string `xml:"valign,attr,omitempty"`

	BeatUnit          []BeatUnit          `xml:"beat-unit,omitempty"`
	BeatUnitDot       []BeatUnitDot       `xml:"beat-unit-dot,omitempty"`
	BeatUnitTied      []BeatUnitTied      `xml:"beat-unit-tied,omitempty"`
	MetronomeArrows   []MetronomeArrows   `xml:"metronome-arrows,omitempty"`
	MetronomeNote     []MetronomeNote     `xml:"metronome-note,omitempty"`
	MetronomeRelation []MetronomeRelation `xml:"metronome-relation,omitempty"`
	PerMinute         []PerMinute         `xml:"per-minute,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Metronome) ToMXML() *MXML {
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
	attributes["justify"] = r.JustifyAttr
	attributes["parentheses"] = r.ParenthesesAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["valign"] = r.ValignAttr

	children["beat-unit"] = make([]*MXML, len(r.BeatUnit))
	for i, c := range r.BeatUnit {
		children["beat-unit"][i] = c.ToMXML()
	}

	children["beat-unit-dot"] = make([]*MXML, len(r.BeatUnitDot))
	for i, c := range r.BeatUnitDot {
		children["beat-unit-dot"][i] = c.ToMXML()
	}

	children["beat-unit-tied"] = make([]*MXML, len(r.BeatUnitTied))
	for i, c := range r.BeatUnitTied {
		children["beat-unit-tied"][i] = c.ToMXML()
	}

	children["metronome-arrows"] = make([]*MXML, len(r.MetronomeArrows))
	for i, c := range r.MetronomeArrows {
		children["metronome-arrows"][i] = c.ToMXML()
	}

	children["metronome-note"] = make([]*MXML, len(r.MetronomeNote))
	for i, c := range r.MetronomeNote {
		children["metronome-note"][i] = c.ToMXML()
	}

	children["metronome-relation"] = make([]*MXML, len(r.MetronomeRelation))
	for i, c := range r.MetronomeRelation {
		children["metronome-relation"][i] = c.ToMXML()
	}

	children["per-minute"] = make([]*MXML, len(r.PerMinute))
	for i, c := range r.PerMinute {
		children["per-minute"][i] = c.ToMXML()
	}

	return newMXML("metronome", strings.TrimSpace(r.IValue), attributes, children)
}
