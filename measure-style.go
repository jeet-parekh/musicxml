package musicxml

import (
	"encoding/xml"
	"strings"
)

// MeasureStyle is the structure for measure-style MusicXML element.
type MeasureStyle struct {
	XMLName xml.Name `xml:"measure-style"`

	ColorAttr      string `xml:"color,attr,omitempty"`
	FontFamilyAttr string `xml:"font-family,attr,omitempty"`
	FontSizeAttr   string `xml:"font-size,attr,omitempty"`
	FontStyleAttr  string `xml:"font-style,attr,omitempty"`
	FontWeightAttr string `xml:"font-weight,attr,omitempty"`
	IdAttr         string `xml:"id,attr,omitempty"`
	NumberAttr     string `xml:"number,attr,omitempty"`

	BeatRepeat    []BeatRepeat    `xml:"beat-repeat,omitempty"`
	MeasureRepeat []MeasureRepeat `xml:"measure-repeat,omitempty"`
	MultipleRest  []MultipleRest  `xml:"multiple-rest,omitempty"`
	Slash         []Slash         `xml:"slash,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *MeasureStyle) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["color"] = r.ColorAttr
	attributes["font-family"] = r.FontFamilyAttr
	attributes["font-size"] = r.FontSizeAttr
	attributes["font-style"] = r.FontStyleAttr
	attributes["font-weight"] = r.FontWeightAttr
	attributes["id"] = r.IdAttr
	attributes["number"] = r.NumberAttr

	children["beat-repeat"] = make([]*MXML, len(r.BeatRepeat))
	for i, c := range r.BeatRepeat {
		children["beat-repeat"][i] = c.ToMXML()
	}

	children["measure-repeat"] = make([]*MXML, len(r.MeasureRepeat))
	for i, c := range r.MeasureRepeat {
		children["measure-repeat"][i] = c.ToMXML()
	}

	children["multiple-rest"] = make([]*MXML, len(r.MultipleRest))
	for i, c := range r.MultipleRest {
		children["multiple-rest"][i] = c.ToMXML()
	}

	children["slash"] = make([]*MXML, len(r.Slash))
	for i, c := range r.Slash {
		children["slash"][i] = c.ToMXML()
	}

	return newMXML("measure-style", strings.TrimSpace(r.IValue), attributes, children)
}
