package musicxml

import (
	"encoding/xml"
	"strings"
)

// Pppppp is the structure for pppppp MusicXML element.
type Pppppp struct {
	XMLName xml.Name `xml:"pppppp"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Pppppp) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("pppppp", strings.TrimSpace(r.IValue), attributes, children)
}
