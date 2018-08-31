package musicxml

import (
	"encoding/xml"
	"strings"
)

// Encoding is the structure for encoding MusicXML element.
type Encoding struct {
	XMLName xml.Name `xml:"encoding"`

	Encoder             []Encoder             `xml:"encoder,omitempty"`
	EncodingDate        []EncodingDate        `xml:"encoding-date,omitempty"`
	EncodingDescription []EncodingDescription `xml:"encoding-description,omitempty"`
	Software            []Software            `xml:"software,omitempty"`
	Supports            []Supports            `xml:"supports,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Encoding) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	children["encoder"] = make([]*MXML, len(r.Encoder))
	for i, c := range r.Encoder {
		children["encoder"][i] = c.ToMXML()
	}

	children["encoding-date"] = make([]*MXML, len(r.EncodingDate))
	for i, c := range r.EncodingDate {
		children["encoding-date"][i] = c.ToMXML()
	}

	children["encoding-description"] = make([]*MXML, len(r.EncodingDescription))
	for i, c := range r.EncodingDescription {
		children["encoding-description"][i] = c.ToMXML()
	}

	children["software"] = make([]*MXML, len(r.Software))
	for i, c := range r.Software {
		children["software"][i] = c.ToMXML()
	}

	children["supports"] = make([]*MXML, len(r.Supports))
	for i, c := range r.Supports {
		children["supports"][i] = c.ToMXML()
	}

	return newMXML("encoding", strings.TrimSpace(r.IValue), attributes, children)
}
