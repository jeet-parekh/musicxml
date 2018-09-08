package mxml

import (
	"encoding/xml"
	"strings"
)

// Measure is the structure for measure MusicXML element.
type Measure struct {
	XMLName xml.Name `xml:"measure"`

	IdAttr             string `xml:"id,attr,omitempty"`
	ImplicitAttr       string `xml:"implicit,attr,omitempty"`
	NonControllingAttr string `xml:"non-controlling,attr,omitempty"`
	NumberAttr         string `xml:"number,attr,omitempty"`
	TextAttr           string `xml:"text,attr,omitempty"`
	WidthAttr          string `xml:"width,attr,omitempty"`

	Attributes  []Attributes  `xml:"attributes,omitempty"`
	Backup      []Backup      `xml:"backup,omitempty"`
	Barline     []Barline     `xml:"barline,omitempty"`
	Bookmark    []Bookmark    `xml:"bookmark,omitempty"`
	Direction   []Direction   `xml:"direction,omitempty"`
	FiguredBass []FiguredBass `xml:"figured-bass,omitempty"`
	Forward     []Forward     `xml:"forward,omitempty"`
	Grouping    []Grouping    `xml:"grouping,omitempty"`
	Harmony     []Harmony     `xml:"harmony,omitempty"`
	Link        []Link        `xml:"link,omitempty"`
	Note        []Note        `xml:"note,omitempty"`
	Part        []Part        `xml:"part,omitempty"`
	Print       []Print       `xml:"print,omitempty"`
	Sound       []Sound       `xml:"sound,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Measure) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["id"] = r.IdAttr
	attributes["implicit"] = r.ImplicitAttr
	attributes["non-controlling"] = r.NonControllingAttr
	attributes["number"] = r.NumberAttr
	attributes["text"] = r.TextAttr
	attributes["width"] = r.WidthAttr

	children["attributes"] = make([]*MXML, len(r.Attributes))
	for i, c := range r.Attributes {
		children["attributes"][i] = c.ToMXML()
	}

	children["backup"] = make([]*MXML, len(r.Backup))
	for i, c := range r.Backup {
		children["backup"][i] = c.ToMXML()
	}

	children["barline"] = make([]*MXML, len(r.Barline))
	for i, c := range r.Barline {
		children["barline"][i] = c.ToMXML()
	}

	children["bookmark"] = make([]*MXML, len(r.Bookmark))
	for i, c := range r.Bookmark {
		children["bookmark"][i] = c.ToMXML()
	}

	children["direction"] = make([]*MXML, len(r.Direction))
	for i, c := range r.Direction {
		children["direction"][i] = c.ToMXML()
	}

	children["figured-bass"] = make([]*MXML, len(r.FiguredBass))
	for i, c := range r.FiguredBass {
		children["figured-bass"][i] = c.ToMXML()
	}

	children["forward"] = make([]*MXML, len(r.Forward))
	for i, c := range r.Forward {
		children["forward"][i] = c.ToMXML()
	}

	children["grouping"] = make([]*MXML, len(r.Grouping))
	for i, c := range r.Grouping {
		children["grouping"][i] = c.ToMXML()
	}

	children["harmony"] = make([]*MXML, len(r.Harmony))
	for i, c := range r.Harmony {
		children["harmony"][i] = c.ToMXML()
	}

	children["link"] = make([]*MXML, len(r.Link))
	for i, c := range r.Link {
		children["link"][i] = c.ToMXML()
	}

	children["note"] = make([]*MXML, len(r.Note))
	for i, c := range r.Note {
		children["note"][i] = c.ToMXML()
	}

	children["part"] = make([]*MXML, len(r.Part))
	for i, c := range r.Part {
		children["part"][i] = c.ToMXML()
	}

	children["print"] = make([]*MXML, len(r.Print))
	for i, c := range r.Print {
		children["print"][i] = c.ToMXML()
	}

	children["sound"] = make([]*MXML, len(r.Sound))
	for i, c := range r.Sound {
		children["sound"][i] = c.ToMXML()
	}

	return newMXML("measure", strings.TrimSpace(r.IValue), attributes, children)
}
