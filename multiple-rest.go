package musicxml

import (
	"encoding/xml"
	"strings"
)

// MultipleRest is the structure for multiple-rest MusicXML element.
type MultipleRest struct {
	XMLName xml.Name `xml:"multiple-rest"`

	UseSymbolsAttr string `xml:"use-symbols,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MultipleRest) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["use-symbols"] = r.UseSymbolsAttr

	return newMXML("multiple-rest", strings.TrimSpace(r.IValue), attributes, children)
}
