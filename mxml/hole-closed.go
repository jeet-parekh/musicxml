package mxml

import (
	"encoding/xml"
	"strings"
)

// HoleClosed is the structure for hole-closed MusicXML element.
type HoleClosed struct {
	XMLName xml.Name `xml:"hole-closed"`

	LocationAttr string `xml:"location,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *HoleClosed) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["location"] = r.LocationAttr

	return newMXML("hole-closed", strings.TrimSpace(r.IValue), attributes, children)
}
