package musicxml

import (
	"encoding/xml"
	"strings"
)

// Lyric is the structure for lyric MusicXML element.
type Lyric struct {
	XMLName xml.Name `xml:"lyric"`

	ColorAttr       string `xml:"color,attr,omitempty"`
	DefaultXAttr    string `xml:"default-x,attr,omitempty"`
	DefaultYAttr    string `xml:"default-y,attr,omitempty"`
	IdAttr          string `xml:"id,attr,omitempty"`
	JustifyAttr     string `xml:"justify,attr,omitempty"`
	NameAttr        string `xml:"name,attr,omitempty"`
	NumberAttr      string `xml:"number,attr,omitempty"`
	PlacementAttr   string `xml:"placement,attr,omitempty"`
	PrintObjectAttr string `xml:"print-object,attr,omitempty"`
	RelativeXAttr   string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr   string `xml:"relative-y,attr,omitempty"`
	TimeOnlyAttr    string `xml:"time-only,attr,omitempty"`

	Elision      []Elision      `xml:"elision,omitempty"`
	EndLine      []EndLine      `xml:"end-line,omitempty"`
	EndParagraph []EndParagraph `xml:"end-paragraph,omitempty"`
	Extend       []Extend       `xml:"extend,omitempty"`
	Footnote     []Footnote     `xml:"footnote,omitempty"`
	Humming      []Humming      `xml:"humming,omitempty"`
	Laughing     []Laughing     `xml:"laughing,omitempty"`
	Level        []Level        `xml:"level,omitempty"`
	Syllabic     []Syllabic     `xml:"syllabic,omitempty"`
	Text         []Text         `xml:"text,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Lyric) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["id"] = r.IdAttr
	attributes["justify"] = r.JustifyAttr
	attributes["name"] = r.NameAttr
	attributes["number"] = r.NumberAttr
	attributes["placement"] = r.PlacementAttr
	attributes["print-object"] = r.PrintObjectAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["time-only"] = r.TimeOnlyAttr

	children["elision"] = make([]*MXML, len(r.Elision))
	for i, c := range r.Elision {
		children["elision"][i] = c.ToMXML()
	}

	children["end-line"] = make([]*MXML, len(r.EndLine))
	for i, c := range r.EndLine {
		children["end-line"][i] = c.ToMXML()
	}

	children["end-paragraph"] = make([]*MXML, len(r.EndParagraph))
	for i, c := range r.EndParagraph {
		children["end-paragraph"][i] = c.ToMXML()
	}

	children["extend"] = make([]*MXML, len(r.Extend))
	for i, c := range r.Extend {
		children["extend"][i] = c.ToMXML()
	}

	children["footnote"] = make([]*MXML, len(r.Footnote))
	for i, c := range r.Footnote {
		children["footnote"][i] = c.ToMXML()
	}

	children["humming"] = make([]*MXML, len(r.Humming))
	for i, c := range r.Humming {
		children["humming"][i] = c.ToMXML()
	}

	children["laughing"] = make([]*MXML, len(r.Laughing))
	for i, c := range r.Laughing {
		children["laughing"][i] = c.ToMXML()
	}

	children["level"] = make([]*MXML, len(r.Level))
	for i, c := range r.Level {
		children["level"][i] = c.ToMXML()
	}

	children["syllabic"] = make([]*MXML, len(r.Syllabic))
	for i, c := range r.Syllabic {
		children["syllabic"][i] = c.ToMXML()
	}

	children["text"] = make([]*MXML, len(r.Text))
	for i, c := range r.Text {
		children["text"][i] = c.ToMXML()
	}

	return newMXML("lyric", strings.TrimSpace(r.IValue), attributes, children)
}
