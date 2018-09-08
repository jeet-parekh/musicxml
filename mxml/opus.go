package mxml

import (
	"encoding/xml"
	"strings"
)

// Opus is the structure for opus MusicXML element.
type Opus struct {
	XMLName xml.Name `xml:"opus"`

	ActuateAttr string `xml:"actuate,attr,omitempty"`
	HrefAttr    string `xml:"href,attr,omitempty"`
	RoleAttr    string `xml:"role,attr,omitempty"`
	ShowAttr    string `xml:"show,attr,omitempty"`
	TitleAttr   string `xml:"title,attr,omitempty"`
	TypeAttr    string `xml:"type,attr,omitempty"`
	VersionAttr string `xml:"version,attr,omitempty"`

	Opus     []Opus     `xml:"opus,omitempty"`
	OpusLink []OpusLink `xml:"opus-link,omitempty"`
	Score    []Score    `xml:"score,omitempty"`
	Title    []Title    `xml:"title,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Opus) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["actuate"] = r.ActuateAttr
	attributes["href"] = r.HrefAttr
	attributes["role"] = r.RoleAttr
	attributes["show"] = r.ShowAttr
	attributes["title"] = r.TitleAttr
	attributes["type"] = r.TypeAttr
	attributes["version"] = r.VersionAttr

	children["opus"] = make([]*MXML, len(r.Opus))
	for i, c := range r.Opus {
		children["opus"][i] = c.ToMXML()
	}

	children["opus-link"] = make([]*MXML, len(r.OpusLink))
	for i, c := range r.OpusLink {
		children["opus-link"][i] = c.ToMXML()
	}

	children["score"] = make([]*MXML, len(r.Score))
	for i, c := range r.Score {
		children["score"][i] = c.ToMXML()
	}

	children["title"] = make([]*MXML, len(r.Title))
	for i, c := range r.Title {
		children["title"][i] = c.ToMXML()
	}

	return newMXML("opus", strings.TrimSpace(r.IValue), attributes, children)
}
