package musicxml

import (
	"encoding/xml"
	"strings"
)

// BeatRepeat is the structure for beat-repeat MusicXML element.
type BeatRepeat struct {
	XMLName xml.Name `xml:"beat-repeat"`

	SlashesAttr string `xml:"slashes,attr,omitempty"`
	TypeAttr    string `xml:"type,attr,omitempty"`
	UseDotsAttr string `xml:"use-dots,attr,omitempty"`

	ExceptVoice []ExceptVoice `xml:"except-voice,omitempty"`
	SlashDot    []SlashDot    `xml:"slash-dot,omitempty"`
	SlashType   []SlashType   `xml:"slash-type,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *BeatRepeat) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["slashes"] = r.SlashesAttr
	attributes["type"] = r.TypeAttr
	attributes["use-dots"] = r.UseDotsAttr

	children["except-voice"] = make([]*MXML, len(r.ExceptVoice))
	for i, c := range r.ExceptVoice {
		children["except-voice"][i] = c.ToMXML()
	}

	children["slash-dot"] = make([]*MXML, len(r.SlashDot))
	for i, c := range r.SlashDot {
		children["slash-dot"][i] = c.ToMXML()
	}

	children["slash-type"] = make([]*MXML, len(r.SlashType))
	for i, c := range r.SlashType {
		children["slash-type"][i] = c.ToMXML()
	}

	return newMXML("beat-repeat", strings.TrimSpace(r.IValue), attributes, children)
}
