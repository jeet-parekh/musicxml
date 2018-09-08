package mxml

import (
	"encoding/xml"
	"strings"
)

// PageLayout is the structure for page-layout MusicXML element.
type PageLayout struct {
	XMLName xml.Name `xml:"page-layout"`

	PageHeight  []PageHeight  `xml:"page-height,omitempty"`
	PageMargins []PageMargins `xml:"page-margins,omitempty"`
	PageWidth   []PageWidth   `xml:"page-width,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *PageLayout) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["page-height"] = make([]*MXML, len(r.PageHeight))
	for i, c := range r.PageHeight {
		children["page-height"][i] = c.ToMXML()
	}

	children["page-margins"] = make([]*MXML, len(r.PageMargins))
	for i, c := range r.PageMargins {
		children["page-margins"][i] = c.ToMXML()
	}

	children["page-width"] = make([]*MXML, len(r.PageWidth))
	for i, c := range r.PageWidth {
		children["page-width"][i] = c.ToMXML()
	}

	return newMXML("page-layout", strings.TrimSpace(r.IValue), attributes, children)
}
