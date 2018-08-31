package musicxml

import (
	"encoding/xml"
	"strings"
)

// Tuplet is the structure for tuplet MusicXML element.
type Tuplet struct {
	XMLName xml.Name `xml:"tuplet"`

	BracketAttr    string `xml:"bracket,attr,omitempty"`
	DefaultXAttr   string `xml:"default-x,attr,omitempty"`
	DefaultYAttr   string `xml:"default-y,attr,omitempty"`
	IdAttr         string `xml:"id,attr,omitempty"`
	LineShapeAttr  string `xml:"line-shape,attr,omitempty"`
	NumberAttr     string `xml:"number,attr,omitempty"`
	PlacementAttr  string `xml:"placement,attr,omitempty"`
	RelativeXAttr  string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr  string `xml:"relative-y,attr,omitempty"`
	ShowNumberAttr string `xml:"show-number,attr,omitempty"`
	ShowTypeAttr   string `xml:"show-type,attr,omitempty"`
	TypeAttr       string `xml:"type,attr,omitempty"`

	TupletActual []TupletActual `xml:"tuplet-actual,omitempty"`
	TupletNormal []TupletNormal `xml:"tuplet-normal,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Tuplet) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["bracket"] = r.BracketAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["id"] = r.IdAttr
	attributes["line-shape"] = r.LineShapeAttr
	attributes["number"] = r.NumberAttr
	attributes["placement"] = r.PlacementAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["show-number"] = r.ShowNumberAttr
	attributes["show-type"] = r.ShowTypeAttr
	attributes["type"] = r.TypeAttr

	children["tuplet-actual"] = make([]*MXML, len(r.TupletActual))
	for i, c := range r.TupletActual {
		children["tuplet-actual"][i] = c.ToMXML()
	}

	children["tuplet-normal"] = make([]*MXML, len(r.TupletNormal))
	for i, c := range r.TupletNormal {
		children["tuplet-normal"][i] = c.ToMXML()
	}

	return newMXML("tuplet", strings.TrimSpace(r.IValue), attributes, children)
}
