package mxml

import (
	"encoding/xml"
	"strings"
)

// Note is the structure for note MusicXML element.
type Note struct {
	XMLName xml.Name `xml:"note"`

	AttackAttr       string `xml:"attack,attr,omitempty"`
	ColorAttr        string `xml:"color,attr,omitempty"`
	DefaultXAttr     string `xml:"default-x,attr,omitempty"`
	DefaultYAttr     string `xml:"default-y,attr,omitempty"`
	DynamicsAttr     string `xml:"dynamics,attr,omitempty"`
	EndDynamicsAttr  string `xml:"end-dynamics,attr,omitempty"`
	FontFamilyAttr   string `xml:"font-family,attr,omitempty"`
	FontSizeAttr     string `xml:"font-size,attr,omitempty"`
	FontStyleAttr    string `xml:"font-style,attr,omitempty"`
	FontWeightAttr   string `xml:"font-weight,attr,omitempty"`
	IdAttr           string `xml:"id,attr,omitempty"`
	PizzicatoAttr    string `xml:"pizzicato,attr,omitempty"`
	PrintDotAttr     string `xml:"print-dot,attr,omitempty"`
	PrintLegerAttr   string `xml:"print-leger,attr,omitempty"`
	PrintLyricAttr   string `xml:"print-lyric,attr,omitempty"`
	PrintObjectAttr  string `xml:"print-object,attr,omitempty"`
	PrintSpacingAttr string `xml:"print-spacing,attr,omitempty"`
	RelativeXAttr    string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr    string `xml:"relative-y,attr,omitempty"`
	ReleaseAttr      string `xml:"release,attr,omitempty"`
	TimeOnlyAttr     string `xml:"time-only,attr,omitempty"`

	Accidental       []Accidental       `xml:"accidental,omitempty"`
	Beam             []Beam             `xml:"beam,omitempty"`
	Chord            []Chord            `xml:"chord,omitempty"`
	Cue              []Cue              `xml:"cue,omitempty"`
	Dot              []Dot              `xml:"dot,omitempty"`
	Duration         []Duration         `xml:"duration,omitempty"`
	Footnote         []Footnote         `xml:"footnote,omitempty"`
	Grace            []Grace            `xml:"grace,omitempty"`
	Instrument       []Instrument       `xml:"instrument,omitempty"`
	Level            []Level            `xml:"level,omitempty"`
	Lyric            []Lyric            `xml:"lyric,omitempty"`
	Notations        []Notations        `xml:"notations,omitempty"`
	Notehead         []Notehead         `xml:"notehead,omitempty"`
	NoteheadText     []NoteheadText     `xml:"notehead-text,omitempty"`
	Pitch            []Pitch            `xml:"pitch,omitempty"`
	Play             []Play             `xml:"play,omitempty"`
	Rest             []Rest             `xml:"rest,omitempty"`
	Staff            []Staff            `xml:"staff,omitempty"`
	Stem             []Stem             `xml:"stem,omitempty"`
	Tie              []Tie              `xml:"tie,omitempty"`
	TimeModification []TimeModification `xml:"time-modification,omitempty"`
	Type             []Type             `xml:"type,omitempty"`
	Unpitched        []Unpitched        `xml:"unpitched,omitempty"`
	Voice            []Voice            `xml:"voice,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Note) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["attack"] = r.AttackAttr
	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["dynamics"] = r.DynamicsAttr
	attributes["end-dynamics"] = r.EndDynamicsAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["id"] = r.IdAttr
	attributes["pizzicato"] = r.PizzicatoAttr
	attributes["print-dot"] = r.PrintDotAttr
	attributes["print-leger"] = r.PrintLegerAttr
	attributes["print-lyric"] = r.PrintLyricAttr
	attributes["print-object"] = r.PrintObjectAttr
	attributes["print-spacing"] = r.PrintSpacingAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["release"] = r.ReleaseAttr
	attributes["time-only"] = r.TimeOnlyAttr

	children["accidental"] = make([]*MXML, len(r.Accidental))
	for i, c := range r.Accidental {
		children["accidental"][i] = c.ToMXML()
	}

	children["beam"] = make([]*MXML, len(r.Beam))
	for i, c := range r.Beam {
		children["beam"][i] = c.ToMXML()
	}

	children["chord"] = make([]*MXML, len(r.Chord))
	for i, c := range r.Chord {
		children["chord"][i] = c.ToMXML()
	}

	children["cue"] = make([]*MXML, len(r.Cue))
	for i, c := range r.Cue {
		children["cue"][i] = c.ToMXML()
	}

	children["dot"] = make([]*MXML, len(r.Dot))
	for i, c := range r.Dot {
		children["dot"][i] = c.ToMXML()
	}

	children["duration"] = make([]*MXML, len(r.Duration))
	for i, c := range r.Duration {
		children["duration"][i] = c.ToMXML()
	}

	children["footnote"] = make([]*MXML, len(r.Footnote))
	for i, c := range r.Footnote {
		children["footnote"][i] = c.ToMXML()
	}

	children["grace"] = make([]*MXML, len(r.Grace))
	for i, c := range r.Grace {
		children["grace"][i] = c.ToMXML()
	}

	children["instrument"] = make([]*MXML, len(r.Instrument))
	for i, c := range r.Instrument {
		children["instrument"][i] = c.ToMXML()
	}

	children["level"] = make([]*MXML, len(r.Level))
	for i, c := range r.Level {
		children["level"][i] = c.ToMXML()
	}

	children["lyric"] = make([]*MXML, len(r.Lyric))
	for i, c := range r.Lyric {
		children["lyric"][i] = c.ToMXML()
	}

	children["notations"] = make([]*MXML, len(r.Notations))
	for i, c := range r.Notations {
		children["notations"][i] = c.ToMXML()
	}

	children["notehead"] = make([]*MXML, len(r.Notehead))
	for i, c := range r.Notehead {
		children["notehead"][i] = c.ToMXML()
	}

	children["notehead-text"] = make([]*MXML, len(r.NoteheadText))
	for i, c := range r.NoteheadText {
		children["notehead-text"][i] = c.ToMXML()
	}

	children["pitch"] = make([]*MXML, len(r.Pitch))
	for i, c := range r.Pitch {
		children["pitch"][i] = c.ToMXML()
	}

	children["play"] = make([]*MXML, len(r.Play))
	for i, c := range r.Play {
		children["play"][i] = c.ToMXML()
	}

	children["rest"] = make([]*MXML, len(r.Rest))
	for i, c := range r.Rest {
		children["rest"][i] = c.ToMXML()
	}

	children["staff"] = make([]*MXML, len(r.Staff))
	for i, c := range r.Staff {
		children["staff"][i] = c.ToMXML()
	}

	children["stem"] = make([]*MXML, len(r.Stem))
	for i, c := range r.Stem {
		children["stem"][i] = c.ToMXML()
	}

	children["tie"] = make([]*MXML, len(r.Tie))
	for i, c := range r.Tie {
		children["tie"][i] = c.ToMXML()
	}

	children["time-modification"] = make([]*MXML, len(r.TimeModification))
	for i, c := range r.TimeModification {
		children["time-modification"][i] = c.ToMXML()
	}

	children["type"] = make([]*MXML, len(r.Type))
	for i, c := range r.Type {
		children["type"][i] = c.ToMXML()
	}

	children["unpitched"] = make([]*MXML, len(r.Unpitched))
	for i, c := range r.Unpitched {
		children["unpitched"][i] = c.ToMXML()
	}

	children["voice"] = make([]*MXML, len(r.Voice))
	for i, c := range r.Voice {
		children["voice"][i] = c.ToMXML()
	}

	return newMXML("note", strings.TrimSpace(r.IValue), attributes, children)
}
