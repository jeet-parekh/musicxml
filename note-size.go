package musicxml

import (
	"encoding/xml"
	"strings"
)

// NoteSize is the structure for note-size MusicXML element.
type NoteSize struct {
	XMLName xml.Name `xml:"note-size"`

	TypeAttr string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *NoteSize) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["type"] = r.TypeAttr

	return newMXML("note-size", strings.TrimSpace(r.IValue), attributes, children)
}
