package mxml

import (
	"encoding/xml"
	"strings"
)

// Text is the structure for text MusicXML element.
type Text struct {
	XMLName xml.Name `xml:"text"`

	ColorAttr         string `xml:"color,attr,omitempty"`
	DirAttr           string `xml:"dir,attr,omitempty"`
	FontFamilyAttr    string `xml:"font-family,attr,omitempty"`
	FontSizeAttr      string `xml:"font-size,attr,omitempty"`
	FontStyleAttr     string `xml:"font-style,attr,omitempty"`
	FontWeightAttr    string `xml:"font-weight,attr,omitempty"`
	LangAttr          string `xml:"lang,attr,omitempty"`
	LetterSpacingAttr string `xml:"letter-spacing,attr,omitempty"`
	LineThroughAttr   string `xml:"line-through,attr,omitempty"`
	OverlineAttr      string `xml:"overline,attr,omitempty"`
	RotationAttr      string `xml:"rotation,attr,omitempty"`
	UnderlineAttr     string `xml:"underline,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Text) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["dir"] = r.DirAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["lang"] = r.LangAttr
	attributes["letter-spacing"] = r.LetterSpacingAttr
	attributes["line-through"] = r.LineThroughAttr
	attributes["overline"] = r.OverlineAttr
	attributes["rotation"] = r.RotationAttr
	attributes["underline"] = r.UnderlineAttr

	return newMXML("text", strings.TrimSpace(r.IValue), attributes, children)
}
