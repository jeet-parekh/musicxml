package musicxml

import (
	"encoding/xml"
	"strings"
)

// AccidentalText is the structure for accidental-text MusicXML element.
type AccidentalText struct {
	XMLName xml.Name `xml:"accidental-text"`

	ColorAttr         string `xml:"color,attr,omitempty"`
	DefaultXAttr      string `xml:"default-x,attr,omitempty"`
	DefaultYAttr      string `xml:"default-y,attr,omitempty"`
	DirAttr           string `xml:"dir,attr,omitempty"`
	EnclosureAttr     string `xml:"enclosure,attr,omitempty"`
	FontFamilyAttr    string `xml:"font-family,attr,omitempty"`
	FontSizeAttr      string `xml:"font-size,attr,omitempty"`
	FontStyleAttr     string `xml:"font-style,attr,omitempty"`
	FontWeightAttr    string `xml:"font-weight,attr,omitempty"`
	HalignAttr        string `xml:"halign,attr,omitempty"`
	JustifyAttr       string `xml:"justify,attr,omitempty"`
	LangAttr          string `xml:"lang,attr,omitempty"`
	LetterSpacingAttr string `xml:"letter-spacing,attr,omitempty"`
	LineHeightAttr    string `xml:"line-height,attr,omitempty"`
	LineThroughAttr   string `xml:"line-through,attr,omitempty"`
	OverlineAttr      string `xml:"overline,attr,omitempty"`
	RelativeXAttr     string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr     string `xml:"relative-y,attr,omitempty"`
	RotationAttr      string `xml:"rotation,attr,omitempty"`
	SmuflAttr         string `xml:"smufl,attr,omitempty"`
	SpaceAttr         string `xml:"space,attr,omitempty"`
	UnderlineAttr     string `xml:"underline,attr,omitempty"`
	ValignAttr        string `xml:"valign,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *AccidentalText) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["dir"] = r.DirAttr
	attributes["enclosure"] = r.EnclosureAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["halign"] = r.HalignAttr
	attributes["justify"] = r.JustifyAttr
	attributes["lang"] = r.LangAttr
	attributes["letter-spacing"] = r.LetterSpacingAttr
	attributes["line-height"] = r.LineHeightAttr
	attributes["line-through"] = r.LineThroughAttr
	attributes["overline"] = r.OverlineAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["rotation"] = r.RotationAttr
	attributes["smufl"] = r.SmuflAttr
	attributes["space"] = r.SpaceAttr
	attributes["underline"] = r.UnderlineAttr
	attributes["valign"] = r.ValignAttr

	return newMXML("accidental-text", strings.TrimSpace(r.IValue), attributes, children)
}
