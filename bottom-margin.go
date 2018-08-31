package musicxml

import (
	"encoding/xml"
	"strings"
)

// BottomMargin is the structure for bottom-margin MusicXML element.
type BottomMargin struct {
	XMLName xml.Name `xml:"bottom-margin"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *BottomMargin) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("bottom-margin", strings.TrimSpace(r.IValue), attributes, children)
}
