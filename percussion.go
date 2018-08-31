package musicxml

import (
	"encoding/xml"
	"strings"
)

// Percussion is the structure for percussion MusicXML element.
type Percussion struct {
	XMLName xml.Name `xml:"percussion"`

	ColorAttr      string `xml:"color,attr,omitempty"`
	DefaultXAttr   string `xml:"default-x,attr,omitempty"`
	DefaultYAttr   string `xml:"default-y,attr,omitempty"`
	EnclosureAttr  string `xml:"enclosure,attr,omitempty"`
	FontFamilyAttr string `xml:"font-family,attr,omitempty"`
	FontSizeAttr   string `xml:"font-size,attr,omitempty"`
	FontStyleAttr  string `xml:"font-style,attr,omitempty"`
	FontWeightAttr string `xml:"font-weight,attr,omitempty"`
	HalignAttr     string `xml:"halign,attr,omitempty"`
	IdAttr         string `xml:"id,attr,omitempty"`
	RelativeXAttr  string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr  string `xml:"relative-y,attr,omitempty"`
	ValignAttr     string `xml:"valign,attr,omitempty"`

	Beater          []Beater          `xml:"beater,omitempty"`
	Effect          []Effect          `xml:"effect,omitempty"`
	Glass           []Glass           `xml:"glass,omitempty"`
	Membrane        []Membrane        `xml:"membrane,omitempty"`
	Metal           []Metal           `xml:"metal,omitempty"`
	OtherPercussion []OtherPercussion `xml:"other-percussion,omitempty"`
	Pitched         []Pitched         `xml:"pitched,omitempty"`
	Stick           []Stick           `xml:"stick,omitempty"`
	StickLocation   []StickLocation   `xml:"stick-location,omitempty"`
	Timpani         []Timpani         `xml:"timpani,omitempty"`
	Wood            []Wood            `xml:"wood,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Percussion) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["enclosure"] = r.EnclosureAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["halign"] = r.HalignAttr
	attributes["id"] = r.IdAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["valign"] = r.ValignAttr

	children["beater"] = make([]*MXML, len(r.Beater))
	for i, c := range r.Beater {
		children["beater"][i] = c.ToMXML()
	}

	children["effect"] = make([]*MXML, len(r.Effect))
	for i, c := range r.Effect {
		children["effect"][i] = c.ToMXML()
	}

	children["glass"] = make([]*MXML, len(r.Glass))
	for i, c := range r.Glass {
		children["glass"][i] = c.ToMXML()
	}

	children["membrane"] = make([]*MXML, len(r.Membrane))
	for i, c := range r.Membrane {
		children["membrane"][i] = c.ToMXML()
	}

	children["metal"] = make([]*MXML, len(r.Metal))
	for i, c := range r.Metal {
		children["metal"][i] = c.ToMXML()
	}

	children["other-percussion"] = make([]*MXML, len(r.OtherPercussion))
	for i, c := range r.OtherPercussion {
		children["other-percussion"][i] = c.ToMXML()
	}

	children["pitched"] = make([]*MXML, len(r.Pitched))
	for i, c := range r.Pitched {
		children["pitched"][i] = c.ToMXML()
	}

	children["stick"] = make([]*MXML, len(r.Stick))
	for i, c := range r.Stick {
		children["stick"][i] = c.ToMXML()
	}

	children["stick-location"] = make([]*MXML, len(r.StickLocation))
	for i, c := range r.StickLocation {
		children["stick-location"][i] = c.ToMXML()
	}

	children["timpani"] = make([]*MXML, len(r.Timpani))
	for i, c := range r.Timpani {
		children["timpani"][i] = c.ToMXML()
	}

	children["wood"] = make([]*MXML, len(r.Wood))
	for i, c := range r.Wood {
		children["wood"][i] = c.ToMXML()
	}

	return newMXML("percussion", strings.TrimSpace(r.IValue), attributes, children)
}
