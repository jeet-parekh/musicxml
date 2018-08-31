package musicxml

import (
	"encoding/xml"
	"strings"
)

// PageMargins is the structure for page-margins MusicXML element.
type PageMargins struct {
	XMLName xml.Name `xml:"page-margins"`

	TypeAttr string `xml:"type,attr,omitempty"`

	BottomMargin []BottomMargin `xml:"bottom-margin,omitempty"`
	LeftMargin   []LeftMargin   `xml:"left-margin,omitempty"`
	RightMargin  []RightMargin  `xml:"right-margin,omitempty"`
	TopMargin    []TopMargin    `xml:"top-margin,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *PageMargins) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["type"] = r.TypeAttr

	children["bottom-margin"] = make([]*MXML, len(r.BottomMargin))
	for i, c := range r.BottomMargin {
		children["bottom-margin"][i] = c.ToMXML()
	}

	children["left-margin"] = make([]*MXML, len(r.LeftMargin))
	for i, c := range r.LeftMargin {
		children["left-margin"][i] = c.ToMXML()
	}

	children["right-margin"] = make([]*MXML, len(r.RightMargin))
	for i, c := range r.RightMargin {
		children["right-margin"][i] = c.ToMXML()
	}

	children["top-margin"] = make([]*MXML, len(r.TopMargin))
	for i, c := range r.TopMargin {
		children["top-margin"][i] = c.ToMXML()
	}

	return newMXML("page-margins", strings.TrimSpace(r.IValue), attributes, children)
}
