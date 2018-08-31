package musicxml

import (
	"encoding/xml"
	"strings"
)

// CreditSymbol is the structure for credit-symbol MusicXML element.
type CreditSymbol struct {
	XMLName xml.Name `xml:"credit-symbol"`

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
	IdAttr            string `xml:"id,attr,omitempty"`
	JustifyAttr       string `xml:"justify,attr,omitempty"`
	LetterSpacingAttr string `xml:"letter-spacing,attr,omitempty"`
	LineHeightAttr    string `xml:"line-height,attr,omitempty"`
	LineThroughAttr   string `xml:"line-through,attr,omitempty"`
	OverlineAttr      string `xml:"overline,attr,omitempty"`
	RelativeXAttr     string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr     string `xml:"relative-y,attr,omitempty"`
	RotationAttr      string `xml:"rotation,attr,omitempty"`
	UnderlineAttr     string `xml:"underline,attr,omitempty"`
	ValignAttr        string `xml:"valign,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *CreditSymbol) ToMXML() *MXML {
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
	attributes["id"] = r.IdAttr
	attributes["justify"] = r.JustifyAttr
	attributes["letter-spacing"] = r.LetterSpacingAttr
	attributes["line-height"] = r.LineHeightAttr
	attributes["line-through"] = r.LineThroughAttr
	attributes["overline"] = r.OverlineAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["rotation"] = r.RotationAttr
	attributes["underline"] = r.UnderlineAttr
	attributes["valign"] = r.ValignAttr

	return newMXML("credit-symbol", strings.TrimSpace(r.IValue), attributes, children)
}
