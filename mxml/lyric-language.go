package mxml

import (
	"encoding/xml"
	"strings"
)

// LyricLanguage is the structure for lyric-language MusicXML element.
type LyricLanguage struct {
	XMLName xml.Name `xml:"lyric-language"`

	LangAttr   string `xml:"lang,attr,omitempty"`
	NameAttr   string `xml:"name,attr,omitempty"`
	NumberAttr string `xml:"number,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *LyricLanguage) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["lang"] = r.LangAttr
	attributes["name"] = r.NameAttr
	attributes["number"] = r.NumberAttr

	return newMXML("lyric-language", strings.TrimSpace(r.IValue), attributes, children)
}
