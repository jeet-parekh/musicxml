package mxml

import (
	"encoding/xml"
	"strings"
)

// DirectionType is the structure for direction-type MusicXML element.
type DirectionType struct {
	XMLName xml.Name `xml:"direction-type"`

	IdAttr string `xml:"id,attr,omitempty"`

	AccordionRegistration []AccordionRegistration `xml:"accordion-registration,omitempty"`
	Bracket               []Bracket               `xml:"bracket,omitempty"`
	Coda                  []Coda                  `xml:"coda,omitempty"`
	Damp                  []Damp                  `xml:"damp,omitempty"`
	DampAll               []DampAll               `xml:"damp-all,omitempty"`
	Dashes                []Dashes                `xml:"dashes,omitempty"`
	Dynamics              []Dynamics              `xml:"dynamics,omitempty"`
	Eyeglasses            []Eyeglasses            `xml:"eyeglasses,omitempty"`
	HarpPedals            []HarpPedals            `xml:"harp-pedals,omitempty"`
	Image                 []Image                 `xml:"image,omitempty"`
	Metronome             []Metronome             `xml:"metronome,omitempty"`
	OctaveShift           []OctaveShift           `xml:"octave-shift,omitempty"`
	OtherDirection        []OtherDirection        `xml:"other-direction,omitempty"`
	Pedal                 []Pedal                 `xml:"pedal,omitempty"`
	Percussion            []Percussion            `xml:"percussion,omitempty"`
	PrincipalVoice        []PrincipalVoice        `xml:"principal-voice,omitempty"`
	Rehearsal             []Rehearsal             `xml:"rehearsal,omitempty"`
	Scordatura            []Scordatura            `xml:"scordatura,omitempty"`
	Segno                 []Segno                 `xml:"segno,omitempty"`
	StaffDivide           []StaffDivide           `xml:"staff-divide,omitempty"`
	StringMute            []StringMute            `xml:"string-mute,omitempty"`
	Symbol                []Symbol                `xml:"symbol,omitempty"`
	Wedge                 []Wedge                 `xml:"wedge,omitempty"`
	Words                 []Words                 `xml:"words,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *DirectionType) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["id"] = r.IdAttr

	children["accordion-registration"] = make([]*MXML, len(r.AccordionRegistration))
	for i, c := range r.AccordionRegistration {
		children["accordion-registration"][i] = c.ToMXML()
	}

	children["bracket"] = make([]*MXML, len(r.Bracket))
	for i, c := range r.Bracket {
		children["bracket"][i] = c.ToMXML()
	}

	children["coda"] = make([]*MXML, len(r.Coda))
	for i, c := range r.Coda {
		children["coda"][i] = c.ToMXML()
	}

	children["damp"] = make([]*MXML, len(r.Damp))
	for i, c := range r.Damp {
		children["damp"][i] = c.ToMXML()
	}

	children["damp-all"] = make([]*MXML, len(r.DampAll))
	for i, c := range r.DampAll {
		children["damp-all"][i] = c.ToMXML()
	}

	children["dashes"] = make([]*MXML, len(r.Dashes))
	for i, c := range r.Dashes {
		children["dashes"][i] = c.ToMXML()
	}

	children["dynamics"] = make([]*MXML, len(r.Dynamics))
	for i, c := range r.Dynamics {
		children["dynamics"][i] = c.ToMXML()
	}

	children["eyeglasses"] = make([]*MXML, len(r.Eyeglasses))
	for i, c := range r.Eyeglasses {
		children["eyeglasses"][i] = c.ToMXML()
	}

	children["harp-pedals"] = make([]*MXML, len(r.HarpPedals))
	for i, c := range r.HarpPedals {
		children["harp-pedals"][i] = c.ToMXML()
	}

	children["image"] = make([]*MXML, len(r.Image))
	for i, c := range r.Image {
		children["image"][i] = c.ToMXML()
	}

	children["metronome"] = make([]*MXML, len(r.Metronome))
	for i, c := range r.Metronome {
		children["metronome"][i] = c.ToMXML()
	}

	children["octave-shift"] = make([]*MXML, len(r.OctaveShift))
	for i, c := range r.OctaveShift {
		children["octave-shift"][i] = c.ToMXML()
	}

	children["other-direction"] = make([]*MXML, len(r.OtherDirection))
	for i, c := range r.OtherDirection {
		children["other-direction"][i] = c.ToMXML()
	}

	children["pedal"] = make([]*MXML, len(r.Pedal))
	for i, c := range r.Pedal {
		children["pedal"][i] = c.ToMXML()
	}

	children["percussion"] = make([]*MXML, len(r.Percussion))
	for i, c := range r.Percussion {
		children["percussion"][i] = c.ToMXML()
	}

	children["principal-voice"] = make([]*MXML, len(r.PrincipalVoice))
	for i, c := range r.PrincipalVoice {
		children["principal-voice"][i] = c.ToMXML()
	}

	children["rehearsal"] = make([]*MXML, len(r.Rehearsal))
	for i, c := range r.Rehearsal {
		children["rehearsal"][i] = c.ToMXML()
	}

	children["scordatura"] = make([]*MXML, len(r.Scordatura))
	for i, c := range r.Scordatura {
		children["scordatura"][i] = c.ToMXML()
	}

	children["segno"] = make([]*MXML, len(r.Segno))
	for i, c := range r.Segno {
		children["segno"][i] = c.ToMXML()
	}

	children["staff-divide"] = make([]*MXML, len(r.StaffDivide))
	for i, c := range r.StaffDivide {
		children["staff-divide"][i] = c.ToMXML()
	}

	children["string-mute"] = make([]*MXML, len(r.StringMute))
	for i, c := range r.StringMute {
		children["string-mute"][i] = c.ToMXML()
	}

	children["symbol"] = make([]*MXML, len(r.Symbol))
	for i, c := range r.Symbol {
		children["symbol"][i] = c.ToMXML()
	}

	children["wedge"] = make([]*MXML, len(r.Wedge))
	for i, c := range r.Wedge {
		children["wedge"][i] = c.ToMXML()
	}

	children["words"] = make([]*MXML, len(r.Words))
	for i, c := range r.Words {
		children["words"][i] = c.ToMXML()
	}

	return newMXML("direction-type", strings.TrimSpace(r.IValue), attributes, children)
}
