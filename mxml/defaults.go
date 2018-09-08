package mxml

import (
	"encoding/xml"
	"strings"
)

// Defaults is the structure for defaults MusicXML element.
type Defaults struct {
	XMLName xml.Name `xml:"defaults"`

	Appearance    []Appearance    `xml:"appearance,omitempty"`
	LyricFont     []LyricFont     `xml:"lyric-font,omitempty"`
	LyricLanguage []LyricLanguage `xml:"lyric-language,omitempty"`
	MusicFont     []MusicFont     `xml:"music-font,omitempty"`
	PageLayout    []PageLayout    `xml:"page-layout,omitempty"`
	Scaling       []Scaling       `xml:"scaling,omitempty"`
	StaffLayout   []StaffLayout   `xml:"staff-layout,omitempty"`
	SystemLayout  []SystemLayout  `xml:"system-layout,omitempty"`
	WordFont      []WordFont      `xml:"word-font,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Defaults) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["appearance"] = make([]*MXML, len(r.Appearance))
	for i, c := range r.Appearance {
		children["appearance"][i] = c.ToMXML()
	}

	children["lyric-font"] = make([]*MXML, len(r.LyricFont))
	for i, c := range r.LyricFont {
		children["lyric-font"][i] = c.ToMXML()
	}

	children["lyric-language"] = make([]*MXML, len(r.LyricLanguage))
	for i, c := range r.LyricLanguage {
		children["lyric-language"][i] = c.ToMXML()
	}

	children["music-font"] = make([]*MXML, len(r.MusicFont))
	for i, c := range r.MusicFont {
		children["music-font"][i] = c.ToMXML()
	}

	children["page-layout"] = make([]*MXML, len(r.PageLayout))
	for i, c := range r.PageLayout {
		children["page-layout"][i] = c.ToMXML()
	}

	children["scaling"] = make([]*MXML, len(r.Scaling))
	for i, c := range r.Scaling {
		children["scaling"][i] = c.ToMXML()
	}

	children["staff-layout"] = make([]*MXML, len(r.StaffLayout))
	for i, c := range r.StaffLayout {
		children["staff-layout"][i] = c.ToMXML()
	}

	children["system-layout"] = make([]*MXML, len(r.SystemLayout))
	for i, c := range r.SystemLayout {
		children["system-layout"][i] = c.ToMXML()
	}

	children["word-font"] = make([]*MXML, len(r.WordFont))
	for i, c := range r.WordFont {
		children["word-font"][i] = c.ToMXML()
	}

	return newMXML("defaults", strings.TrimSpace(r.IValue), attributes, children)
}
