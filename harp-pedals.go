package musicxml

import (
	"encoding/xml"
	"strings"
)

// HarpPedals is the structure for harp-pedals MusicXML element.
type HarpPedals struct {
	XMLName xml.Name `xml:"harp-pedals"`

	ColorAttr      string `xml:"color,attr,omitempty"`
	DefaultXAttr   string `xml:"default-x,attr,omitempty"`
	DefaultYAttr   string `xml:"default-y,attr,omitempty"`
	FontFamilyAttr string `xml:"font-family,attr,omitempty"`
	FontSizeAttr   string `xml:"font-size,attr,omitempty"`
	FontStyleAttr  string `xml:"font-style,attr,omitempty"`
	FontWeightAttr string `xml:"font-weight,attr,omitempty"`
	HalignAttr     string `xml:"halign,attr,omitempty"`
	IdAttr         string `xml:"id,attr,omitempty"`
	RelativeXAttr  string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr  string `xml:"relative-y,attr,omitempty"`
	ValignAttr     string `xml:"valign,attr,omitempty"`

	PedalTuning []PedalTuning `xml:"pedal-tuning,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *HarpPedals) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["halign"] = r.HalignAttr
	attributes["id"] = r.IdAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["valign"] = r.ValignAttr

	children["pedal-tuning"] = make([]*MXML, len(r.PedalTuning))
	for i, c := range r.PedalTuning {
		children["pedal-tuning"][i] = c.ToMXML()
	}

	return newMXML("harp-pedals", strings.TrimSpace(r.IValue), attributes, children)
}
