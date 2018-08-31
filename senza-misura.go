package musicxml

import (
	"encoding/xml"
	"strings"
)

// SenzaMisura is the structure for senza-misura MusicXML element.
type SenzaMisura struct {
	XMLName xml.Name `xml:"senza-misura"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *SenzaMisura) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("senza-misura", strings.TrimSpace(r.IValue), attributes, children)
}
