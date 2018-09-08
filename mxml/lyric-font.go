package mxml

import (
	"encoding/xml"
	"strings"
)

// LyricFont is the structure for lyric-font MusicXML element.
type LyricFont struct {
	XMLName xml.Name `xml:"lyric-font"`

	FontFamilyAttr string `xml:"font-family,attr,omitempty"`
	FontSizeAttr   string `xml:"font-size,attr,omitempty"`
	FontStyleAttr  string `xml:"font-style,attr,omitempty"`
	FontWeightAttr string `xml:"font-weight,attr,omitempty"`
	NameAttr       string `xml:"name,attr,omitempty"`
	NumberAttr     string `xml:"number,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *LyricFont) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["name"] = r.NameAttr
	attributes["number"] = r.NumberAttr

	return newMXML("lyric-font", strings.TrimSpace(r.IValue), attributes, children)
}
