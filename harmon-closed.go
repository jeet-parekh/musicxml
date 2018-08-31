package musicxml

import (
	"encoding/xml"
	"strings"
)

// HarmonClosed is the structure for harmon-closed MusicXML element.
type HarmonClosed struct {
	XMLName xml.Name `xml:"harmon-closed"`

	LocationAttr string `xml:"location,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *HarmonClosed) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["location"] = r.LocationAttr

	return newMXML("harmon-closed", strings.TrimSpace(r.IValue), attributes, children)
}
