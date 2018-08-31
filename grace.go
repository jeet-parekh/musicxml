package musicxml

import (
	"encoding/xml"
	"strings"
)

// Grace is the structure for grace MusicXML element.
type Grace struct {
	XMLName xml.Name `xml:"grace"`

	MakeTimeAttr           string `xml:"make-time,attr,omitempty"`
	SlashAttr              string `xml:"slash,attr,omitempty"`
	StealTimeFollowingAttr string `xml:"steal-time-following,attr,omitempty"`
	StealTimePreviousAttr  string `xml:"steal-time-previous,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Grace) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["make-time"] = r.MakeTimeAttr
	attributes["slash"] = r.SlashAttr
	attributes["steal-time-following"] = r.StealTimeFollowingAttr
	attributes["steal-time-previous"] = r.StealTimePreviousAttr

	return newMXML("grace", strings.TrimSpace(r.IValue), attributes, children)
}
