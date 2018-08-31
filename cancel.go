package musicxml

import (
	"encoding/xml"
	"strings"
)

// Cancel is the structure for cancel MusicXML element.
type Cancel struct {
	XMLName xml.Name `xml:"cancel"`

	LocationAttr string `xml:"location,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Cancel) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["location"] = r.LocationAttr

	return newMXML("cancel", strings.TrimSpace(r.IValue), attributes, children)
}
