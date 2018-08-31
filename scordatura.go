package musicxml

import (
	"encoding/xml"
	"strings"
)

// Scordatura is the structure for scordatura MusicXML element.
type Scordatura struct {
	XMLName xml.Name `xml:"scordatura"`

	IdAttr string `xml:"id,attr,omitempty"`

	Accord []Accord `xml:"accord,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Scordatura) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["id"] = r.IdAttr

	children["accord"] = make([]*MXML, len(r.Accord))
	for i, c := range r.Accord {
		children["accord"][i] = c.ToMXML()
	}

	return newMXML("scordatura", strings.TrimSpace(r.IValue), attributes, children)
}
