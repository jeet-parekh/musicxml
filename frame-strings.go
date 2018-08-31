package musicxml

import (
	"encoding/xml"
	"strings"
)

// FrameStrings is the structure for frame-strings MusicXML element.
type FrameStrings struct {
	XMLName xml.Name `xml:"frame-strings"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *FrameStrings) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("frame-strings", strings.TrimSpace(r.IValue), attributes, children)
}
