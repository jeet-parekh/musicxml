package musicxml

import (
	"encoding/xml"
	"strings"
)

// Clef is the structure for clef MusicXML element.
type Clef struct {
	XMLName xml.Name `xml:"clef"`

	AdditionalAttr   string `xml:"additional,attr,omitempty"`
	AfterBarlineAttr string `xml:"after-barline,attr,omitempty"`
	ColorAttr        string `xml:"color,attr,omitempty"`
	DefaultXAttr     string `xml:"default-x,attr,omitempty"`
	DefaultYAttr     string `xml:"default-y,attr,omitempty"`
	FontFamilyAttr   string `xml:"font-family,attr,omitempty"`
	FontSizeAttr     string `xml:"font-size,attr,omitempty"`
	FontStyleAttr    string `xml:"font-style,attr,omitempty"`
	FontWeightAttr   string `xml:"font-weight,attr,omitempty"`
	IdAttr           string `xml:"id,attr,omitempty"`
	NumberAttr       string `xml:"number,attr,omitempty"`
	PrintObjectAttr  string `xml:"print-object,attr,omitempty"`
	RelativeXAttr    string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr    string `xml:"relative-y,attr,omitempty"`
	SizeAttr         string `xml:"size,attr,omitempty"`

	ClefOctaveChange []ClefOctaveChange `xml:"clef-octave-change,omitempty"`
	Line             []Line             `xml:"line,omitempty"`
	Sign             []Sign             `xml:"sign,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Clef) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["additional"] = r.AdditionalAttr
	attributes["after-barline"] = r.AfterBarlineAttr
	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["id"] = r.IdAttr
	attributes["number"] = r.NumberAttr
	attributes["print-object"] = r.PrintObjectAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["size"] = r.SizeAttr

	children["clef-octave-change"] = make([]*MXML, len(r.ClefOctaveChange))
	for i, c := range r.ClefOctaveChange {
		children["clef-octave-change"][i] = c.ToMXML()
	}

	children["line"] = make([]*MXML, len(r.Line))
	for i, c := range r.Line {
		children["line"][i] = c.ToMXML()
	}

	children["sign"] = make([]*MXML, len(r.Sign))
	for i, c := range r.Sign {
		children["sign"][i] = c.ToMXML()
	}

	return newMXML("clef", strings.TrimSpace(r.IValue), attributes, children)
}
