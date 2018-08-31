package musicxml

import (
	"encoding/xml"
	"strings"
)

// Degree is the structure for degree MusicXML element.
type Degree struct {
	XMLName xml.Name `xml:"degree"`

	PrintObjectAttr string `xml:"print-object,attr,omitempty"`

	DegreeAlter []DegreeAlter `xml:"degree-alter,omitempty"`
	DegreeType  []DegreeType  `xml:"degree-type,omitempty"`
	DegreeValue []DegreeValue `xml:"degree-value,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Degree) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["print-object"] = r.PrintObjectAttr

	children["degree-alter"] = make([]*MXML, len(r.DegreeAlter))
	for i, c := range r.DegreeAlter {
		children["degree-alter"][i] = c.ToMXML()
	}

	children["degree-type"] = make([]*MXML, len(r.DegreeType))
	for i, c := range r.DegreeType {
		children["degree-type"][i] = c.ToMXML()
	}

	children["degree-value"] = make([]*MXML, len(r.DegreeValue))
	for i, c := range r.DegreeValue {
		children["degree-value"][i] = c.ToMXML()
	}

	return newMXML("degree", strings.TrimSpace(r.IValue), attributes, children)
}
