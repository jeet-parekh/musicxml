package musicxml

import (
	"encoding/xml"
	"strings"
)

// Tied is the structure for tied MusicXML element.
type Tied struct {
	XMLName xml.Name `xml:"tied"`

	BezierOffsetAttr  string `xml:"bezier-offset,attr,omitempty"`
	BezierOffset2Attr string `xml:"bezier-offset2,attr,omitempty"`
	BezierXAttr       string `xml:"bezier-x,attr,omitempty"`
	BezierX2Attr      string `xml:"bezier-x2,attr,omitempty"`
	BezierYAttr       string `xml:"bezier-y,attr,omitempty"`
	BezierY2Attr      string `xml:"bezier-y2,attr,omitempty"`
	ColorAttr         string `xml:"color,attr,omitempty"`
	DashLengthAttr    string `xml:"dash-length,attr,omitempty"`
	DefaultXAttr      string `xml:"default-x,attr,omitempty"`
	DefaultYAttr      string `xml:"default-y,attr,omitempty"`
	IdAttr            string `xml:"id,attr,omitempty"`
	LineTypeAttr      string `xml:"line-type,attr,omitempty"`
	NumberAttr        string `xml:"number,attr,omitempty"`
	OrientationAttr   string `xml:"orientation,attr,omitempty"`
	PlacementAttr     string `xml:"placement,attr,omitempty"`
	RelativeXAttr     string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr     string `xml:"relative-y,attr,omitempty"`
	SpaceLengthAttr   string `xml:"space-length,attr,omitempty"`
	TypeAttr          string `xml:"type,attr,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Tied) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["bezier-offset"] = r.BezierOffsetAttr
	attributes["bezier-offset2"] = r.BezierOffset2Attr
	attributes["bezier-x"] = r.BezierXAttr
	attributes["bezier-x2"] = r.BezierX2Attr
	attributes["bezier-y"] = r.BezierYAttr
	attributes["bezier-y2"] = r.BezierY2Attr
	attributes["color"] = r.ColorAttr
	attributes["dash-length"] = r.DashLengthAttr
	attributes["default-x"] = r.DefaultXAttr
	attributes["default-y"] = r.DefaultYAttr
	attributes["id"] = r.IdAttr
	attributes["line-type"] = r.LineTypeAttr
	attributes["number"] = r.NumberAttr
	attributes["orientation"] = r.OrientationAttr
	attributes["placement"] = r.PlacementAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["space-length"] = r.SpaceLengthAttr
	attributes["type"] = r.TypeAttr

	return newMXML("tied", strings.TrimSpace(r.IValue), attributes, children)
}
