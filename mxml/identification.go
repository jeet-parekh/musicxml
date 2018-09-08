package mxml

import (
	"encoding/xml"
	"strings"
)

// Identification is the structure for identification MusicXML element.
type Identification struct {
	XMLName xml.Name `xml:"identification"`

	Creator       []Creator       `xml:"creator,omitempty"`
	Encoding      []Encoding      `xml:"encoding,omitempty"`
	Miscellaneous []Miscellaneous `xml:"miscellaneous,omitempty"`
	Relation      []Relation      `xml:"relation,omitempty"`
	Rights        []Rights        `xml:"rights,omitempty"`
	Source        []Source        `xml:"source,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Identification) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["creator"] = make([]*MXML, len(r.Creator))
	for i, c := range r.Creator {
		children["creator"][i] = c.ToMXML()
	}

	children["encoding"] = make([]*MXML, len(r.Encoding))
	for i, c := range r.Encoding {
		children["encoding"][i] = c.ToMXML()
	}

	children["miscellaneous"] = make([]*MXML, len(r.Miscellaneous))
	for i, c := range r.Miscellaneous {
		children["miscellaneous"][i] = c.ToMXML()
	}

	children["relation"] = make([]*MXML, len(r.Relation))
	for i, c := range r.Relation {
		children["relation"][i] = c.ToMXML()
	}

	children["rights"] = make([]*MXML, len(r.Rights))
	for i, c := range r.Rights {
		children["rights"][i] = c.ToMXML()
	}

	children["source"] = make([]*MXML, len(r.Source))
	for i, c := range r.Source {
		children["source"][i] = c.ToMXML()
	}

	return newMXML("identification", strings.TrimSpace(r.IValue), attributes, children)
}
