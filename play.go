package musicxml

import (
	"encoding/xml"
	"strings"
)

// Play is the structure for play MusicXML element.
type Play struct {
	XMLName xml.Name `xml:"play"`

	IdAttr string `xml:"id,attr,omitempty"`

	Ipa         []Ipa         `xml:"ipa,omitempty"`
	Mute        []Mute        `xml:"mute,omitempty"`
	OtherPlay   []OtherPlay   `xml:"other-play,omitempty"`
	SemiPitched []SemiPitched `xml:"semi-pitched,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Play) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["id"] = r.IdAttr

	children["ipa"] = make([]*MXML, len(r.Ipa))
	for i, c := range r.Ipa {
		children["ipa"][i] = c.ToMXML()
	}

	children["mute"] = make([]*MXML, len(r.Mute))
	for i, c := range r.Mute {
		children["mute"][i] = c.ToMXML()
	}

	children["other-play"] = make([]*MXML, len(r.OtherPlay))
	for i, c := range r.OtherPlay {
		children["other-play"][i] = c.ToMXML()
	}

	children["semi-pitched"] = make([]*MXML, len(r.SemiPitched))
	for i, c := range r.SemiPitched {
		children["semi-pitched"][i] = c.ToMXML()
	}

	return newMXML("play", strings.TrimSpace(r.IValue), attributes, children)
}
