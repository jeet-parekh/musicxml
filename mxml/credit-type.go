package mxml

import (
	"encoding/xml"
	"strings"
)

// CreditType is the structure for credit-type MusicXML element.
type CreditType struct {
	XMLName xml.Name `xml:"credit-type"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *CreditType) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	return newMXML("credit-type", strings.TrimSpace(r.IValue), attributes, children)
}
