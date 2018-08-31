package musicxml

import (
	"encoding/xml"
	"strings"
)

// Key is the structure for key MusicXML element.
type Key struct {
	XMLName xml.Name `xml:"key"`

	ColorAttr       string `xml:"color,attr,omitempty"`
	DefaultXAttr    string `xml:"default-x,attr,omitempty"`
	DefaultYAttr    string `xml:"default-y,attr,omitempty"`
	FontFamilyAttr  string `xml:"font-family,attr,omitempty"`
	FontSizeAttr    string `xml:"font-size,attr,omitempty"`
	FontStyleAttr   string `xml:"font-style,attr,omitempty"`
	FontWeightAttr  string `xml:"font-weight,attr,omitempty"`
	IdAttr          string `xml:"id,attr,omitempty"`
	NumberAttr      string `xml:"number,attr,omitempty"`
	PrintObjectAttr string `xml:"print-object,attr,omitempty"`
	RelativeXAttr   string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr   string `xml:"relative-y,attr,omitempty"`

	Cancel        []Cancel        `xml:"cancel,omitempty"`
	Fifths        []Fifths        `xml:"fifths,omitempty"`
	KeyAccidental []KeyAccidental `xml:"key-accidental,omitempty"`
	KeyAlter      []KeyAlter      `xml:"key-alter,omitempty"`
	KeyOctave     []KeyOctave     `xml:"key-octave,omitempty"`
	KeyStep       []KeyStep       `xml:"key-step,omitempty"`
	Mode          []Mode          `xml:"mode,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Key) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

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

	children["cancel"] = make([]*MXML, len(r.Cancel))
	for i, c := range r.Cancel {
		children["cancel"][i] = c.ToMXML()
	}

	children["fifths"] = make([]*MXML, len(r.Fifths))
	for i, c := range r.Fifths {
		children["fifths"][i] = c.ToMXML()
	}

	children["key-accidental"] = make([]*MXML, len(r.KeyAccidental))
	for i, c := range r.KeyAccidental {
		children["key-accidental"][i] = c.ToMXML()
	}

	children["key-alter"] = make([]*MXML, len(r.KeyAlter))
	for i, c := range r.KeyAlter {
		children["key-alter"][i] = c.ToMXML()
	}

	children["key-octave"] = make([]*MXML, len(r.KeyOctave))
	for i, c := range r.KeyOctave {
		children["key-octave"][i] = c.ToMXML()
	}

	children["key-step"] = make([]*MXML, len(r.KeyStep))
	for i, c := range r.KeyStep {
		children["key-step"][i] = c.ToMXML()
	}

	children["mode"] = make([]*MXML, len(r.Mode))
	for i, c := range r.Mode {
		children["mode"][i] = c.ToMXML()
	}

	return newMXML("key", strings.TrimSpace(r.IValue), attributes, children)
}
