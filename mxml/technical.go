package mxml

import (
	"encoding/xml"
	"strings"
)

// Technical is the structure for technical MusicXML element.
type Technical struct {
	XMLName xml.Name `xml:"technical"`

	IdAttr string `xml:"id,attr,omitempty"`

	Arrow          []Arrow          `xml:"arrow,omitempty"`
	Bend           []Bend           `xml:"bend,omitempty"`
	BrassBend      []BrassBend      `xml:"brass-bend,omitempty"`
	DoubleTongue   []DoubleTongue   `xml:"double-tongue,omitempty"`
	DownBow        []DownBow        `xml:"down-bow,omitempty"`
	Fingering      []Fingering      `xml:"fingering,omitempty"`
	Fingernails    []Fingernails    `xml:"fingernails,omitempty"`
	Flip           []Flip           `xml:"flip,omitempty"`
	Fret           []Fret           `xml:"fret,omitempty"`
	Golpe          []Golpe          `xml:"golpe,omitempty"`
	HalfMuted      []HalfMuted      `xml:"half-muted,omitempty"`
	HammerOn       []HammerOn       `xml:"hammer-on,omitempty"`
	Handbell       []Handbell       `xml:"handbell,omitempty"`
	HarmonMute     []HarmonMute     `xml:"harmon-mute,omitempty"`
	Harmonic       []Harmonic       `xml:"harmonic,omitempty"`
	Heel           []Heel           `xml:"heel,omitempty"`
	Hole           []Hole           `xml:"hole,omitempty"`
	Open           []Open           `xml:"open,omitempty"`
	OpenString     []OpenString     `xml:"open-string,omitempty"`
	OtherTechnical []OtherTechnical `xml:"other-technical,omitempty"`
	Pluck          []Pluck          `xml:"pluck,omitempty"`
	PullOff        []PullOff        `xml:"pull-off,omitempty"`
	Smear          []Smear          `xml:"smear,omitempty"`
	SnapPizzicato  []SnapPizzicato  `xml:"snap-pizzicato,omitempty"`
	Stopped        []Stopped        `xml:"stopped,omitempty"`
	String         []String         `xml:"string,omitempty"`
	Tap            []Tap            `xml:"tap,omitempty"`
	ThumbPosition  []ThumbPosition  `xml:"thumb-position,omitempty"`
	Toe            []Toe            `xml:"toe,omitempty"`
	TripleTongue   []TripleTongue   `xml:"triple-tongue,omitempty"`
	UpBow          []UpBow          `xml:"up-bow,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Technical) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["id"] = r.IdAttr

	children["arrow"] = make([]*MXML, len(r.Arrow))
	for i, c := range r.Arrow {
		children["arrow"][i] = c.ToMXML()
	}

	children["bend"] = make([]*MXML, len(r.Bend))
	for i, c := range r.Bend {
		children["bend"][i] = c.ToMXML()
	}

	children["brass-bend"] = make([]*MXML, len(r.BrassBend))
	for i, c := range r.BrassBend {
		children["brass-bend"][i] = c.ToMXML()
	}

	children["double-tongue"] = make([]*MXML, len(r.DoubleTongue))
	for i, c := range r.DoubleTongue {
		children["double-tongue"][i] = c.ToMXML()
	}

	children["down-bow"] = make([]*MXML, len(r.DownBow))
	for i, c := range r.DownBow {
		children["down-bow"][i] = c.ToMXML()
	}

	children["fingering"] = make([]*MXML, len(r.Fingering))
	for i, c := range r.Fingering {
		children["fingering"][i] = c.ToMXML()
	}

	children["fingernails"] = make([]*MXML, len(r.Fingernails))
	for i, c := range r.Fingernails {
		children["fingernails"][i] = c.ToMXML()
	}

	children["flip"] = make([]*MXML, len(r.Flip))
	for i, c := range r.Flip {
		children["flip"][i] = c.ToMXML()
	}

	children["fret"] = make([]*MXML, len(r.Fret))
	for i, c := range r.Fret {
		children["fret"][i] = c.ToMXML()
	}

	children["golpe"] = make([]*MXML, len(r.Golpe))
	for i, c := range r.Golpe {
		children["golpe"][i] = c.ToMXML()
	}

	children["half-muted"] = make([]*MXML, len(r.HalfMuted))
	for i, c := range r.HalfMuted {
		children["half-muted"][i] = c.ToMXML()
	}

	children["hammer-on"] = make([]*MXML, len(r.HammerOn))
	for i, c := range r.HammerOn {
		children["hammer-on"][i] = c.ToMXML()
	}

	children["handbell"] = make([]*MXML, len(r.Handbell))
	for i, c := range r.Handbell {
		children["handbell"][i] = c.ToMXML()
	}

	children["harmon-mute"] = make([]*MXML, len(r.HarmonMute))
	for i, c := range r.HarmonMute {
		children["harmon-mute"][i] = c.ToMXML()
	}

	children["harmonic"] = make([]*MXML, len(r.Harmonic))
	for i, c := range r.Harmonic {
		children["harmonic"][i] = c.ToMXML()
	}

	children["heel"] = make([]*MXML, len(r.Heel))
	for i, c := range r.Heel {
		children["heel"][i] = c.ToMXML()
	}

	children["hole"] = make([]*MXML, len(r.Hole))
	for i, c := range r.Hole {
		children["hole"][i] = c.ToMXML()
	}

	children["open"] = make([]*MXML, len(r.Open))
	for i, c := range r.Open {
		children["open"][i] = c.ToMXML()
	}

	children["open-string"] = make([]*MXML, len(r.OpenString))
	for i, c := range r.OpenString {
		children["open-string"][i] = c.ToMXML()
	}

	children["other-technical"] = make([]*MXML, len(r.OtherTechnical))
	for i, c := range r.OtherTechnical {
		children["other-technical"][i] = c.ToMXML()
	}

	children["pluck"] = make([]*MXML, len(r.Pluck))
	for i, c := range r.Pluck {
		children["pluck"][i] = c.ToMXML()
	}

	children["pull-off"] = make([]*MXML, len(r.PullOff))
	for i, c := range r.PullOff {
		children["pull-off"][i] = c.ToMXML()
	}

	children["smear"] = make([]*MXML, len(r.Smear))
	for i, c := range r.Smear {
		children["smear"][i] = c.ToMXML()
	}

	children["snap-pizzicato"] = make([]*MXML, len(r.SnapPizzicato))
	for i, c := range r.SnapPizzicato {
		children["snap-pizzicato"][i] = c.ToMXML()
	}

	children["stopped"] = make([]*MXML, len(r.Stopped))
	for i, c := range r.Stopped {
		children["stopped"][i] = c.ToMXML()
	}

	children["string"] = make([]*MXML, len(r.String))
	for i, c := range r.String {
		children["string"][i] = c.ToMXML()
	}

	children["tap"] = make([]*MXML, len(r.Tap))
	for i, c := range r.Tap {
		children["tap"][i] = c.ToMXML()
	}

	children["thumb-position"] = make([]*MXML, len(r.ThumbPosition))
	for i, c := range r.ThumbPosition {
		children["thumb-position"][i] = c.ToMXML()
	}

	children["toe"] = make([]*MXML, len(r.Toe))
	for i, c := range r.Toe {
		children["toe"][i] = c.ToMXML()
	}

	children["triple-tongue"] = make([]*MXML, len(r.TripleTongue))
	for i, c := range r.TripleTongue {
		children["triple-tongue"][i] = c.ToMXML()
	}

	children["up-bow"] = make([]*MXML, len(r.UpBow))
	for i, c := range r.UpBow {
		children["up-bow"][i] = c.ToMXML()
	}

	return newMXML("technical", strings.TrimSpace(r.IValue), attributes, children)
}
