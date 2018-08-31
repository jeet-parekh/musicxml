package musicxml

import (
	"encoding/xml"
	"strings"
)

// Dynamics is the structure for dynamics MusicXML element.
type Dynamics struct {
	XMLName xml.Name `xml:"dynamics"`

	ColorAttr       string `xml:"color,attr,omitempty"`
	DefaultXAttr    string `xml:"default-x,attr,omitempty"`
	DefaultYAttr    string `xml:"default-y,attr,omitempty"`
	EnclosureAttr   string `xml:"enclosure,attr,omitempty"`
	FontFamilyAttr  string `xml:"font-family,attr,omitempty"`
	FontSizeAttr    string `xml:"font-size,attr,omitempty"`
	FontStyleAttr   string `xml:"font-style,attr,omitempty"`
	FontWeightAttr  string `xml:"font-weight,attr,omitempty"`
	HalignAttr      string `xml:"halign,attr,omitempty"`
	IdAttr          string `xml:"id,attr,omitempty"`
	LineThroughAttr string `xml:"line-through,attr,omitempty"`
	OverlineAttr    string `xml:"overline,attr,omitempty"`
	PlacementAttr   string `xml:"placement,attr,omitempty"`
	RelativeXAttr   string `xml:"relative-x,attr,omitempty"`
	RelativeYAttr   string `xml:"relative-y,attr,omitempty"`
	UnderlineAttr   string `xml:"underline,attr,omitempty"`
	ValignAttr      string `xml:"valign,attr,omitempty"`

	F             []F             `xml:"f,omitempty"`
	Ff            []Ff            `xml:"ff,omitempty"`
	Fff           []Fff           `xml:"fff,omitempty"`
	Ffff          []Ffff          `xml:"ffff,omitempty"`
	Fffff         []Fffff         `xml:"fffff,omitempty"`
	Ffffff        []Ffffff        `xml:"ffffff,omitempty"`
	Fp            []Fp            `xml:"fp,omitempty"`
	Fz            []Fz            `xml:"fz,omitempty"`
	Mf            []Mf            `xml:"mf,omitempty"`
	Mp            []Mp            `xml:"mp,omitempty"`
	N             []N             `xml:"n,omitempty"`
	OtherDynamics []OtherDynamics `xml:"other-dynamics,omitempty"`
	P             []P             `xml:"p,omitempty"`
	Pf            []Pf            `xml:"pf,omitempty"`
	Pp            []Pp            `xml:"pp,omitempty"`
	Ppp           []Ppp           `xml:"ppp,omitempty"`
	Pppp          []Pppp          `xml:"pppp,omitempty"`
	Ppppp         []Ppppp         `xml:"ppppp,omitempty"`
	Pppppp        []Pppppp        `xml:"pppppp,omitempty"`
	Rf            []Rf            `xml:"rf,omitempty"`
	Rfz           []Rfz           `xml:"rfz,omitempty"`
	Sf            []Sf            `xml:"sf,omitempty"`
	Sffz          []Sffz          `xml:"sffz,omitempty"`
	Sfp           []Sfp           `xml:"sfp,omitempty"`
	Sfpp          []Sfpp          `xml:"sfpp,omitempty"`
	Sfz           []Sfz           `xml:"sfz,omitempty"`
	Sfzp          []Sfzp          `xml:"sfzp,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Dynamics) ToMXML() *MXML {
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
	attributes["line-through"] = r.LineThroughAttr
	attributes["overline"] = r.OverlineAttr
	attributes["placement"] = r.PlacementAttr
	attributes["relative-x"] = r.RelativeXAttr
	attributes["relative-y"] = r.RelativeYAttr
	attributes["underline"] = r.UnderlineAttr
	attributes["valign"] = r.ValignAttr

	children["f"] = make([]*MXML, len(r.F))
	for i, c := range r.F {
		children["f"][i] = c.ToMXML()
	}

	children["ff"] = make([]*MXML, len(r.Ff))
	for i, c := range r.Ff {
		children["ff"][i] = c.ToMXML()
	}

	children["fff"] = make([]*MXML, len(r.Fff))
	for i, c := range r.Fff {
		children["fff"][i] = c.ToMXML()
	}

	children["ffff"] = make([]*MXML, len(r.Ffff))
	for i, c := range r.Ffff {
		children["ffff"][i] = c.ToMXML()
	}

	children["fffff"] = make([]*MXML, len(r.Fffff))
	for i, c := range r.Fffff {
		children["fffff"][i] = c.ToMXML()
	}

	children["ffffff"] = make([]*MXML, len(r.Ffffff))
	for i, c := range r.Ffffff {
		children["ffffff"][i] = c.ToMXML()
	}

	children["fp"] = make([]*MXML, len(r.Fp))
	for i, c := range r.Fp {
		children["fp"][i] = c.ToMXML()
	}

	children["fz"] = make([]*MXML, len(r.Fz))
	for i, c := range r.Fz {
		children["fz"][i] = c.ToMXML()
	}

	children["mf"] = make([]*MXML, len(r.Mf))
	for i, c := range r.Mf {
		children["mf"][i] = c.ToMXML()
	}

	children["mp"] = make([]*MXML, len(r.Mp))
	for i, c := range r.Mp {
		children["mp"][i] = c.ToMXML()
	}

	children["n"] = make([]*MXML, len(r.N))
	for i, c := range r.N {
		children["n"][i] = c.ToMXML()
	}

	children["other-dynamics"] = make([]*MXML, len(r.OtherDynamics))
	for i, c := range r.OtherDynamics {
		children["other-dynamics"][i] = c.ToMXML()
	}

	children["p"] = make([]*MXML, len(r.P))
	for i, c := range r.P {
		children["p"][i] = c.ToMXML()
	}

	children["pf"] = make([]*MXML, len(r.Pf))
	for i, c := range r.Pf {
		children["pf"][i] = c.ToMXML()
	}

	children["pp"] = make([]*MXML, len(r.Pp))
	for i, c := range r.Pp {
		children["pp"][i] = c.ToMXML()
	}

	children["ppp"] = make([]*MXML, len(r.Ppp))
	for i, c := range r.Ppp {
		children["ppp"][i] = c.ToMXML()
	}

	children["pppp"] = make([]*MXML, len(r.Pppp))
	for i, c := range r.Pppp {
		children["pppp"][i] = c.ToMXML()
	}

	children["ppppp"] = make([]*MXML, len(r.Ppppp))
	for i, c := range r.Ppppp {
		children["ppppp"][i] = c.ToMXML()
	}

	children["pppppp"] = make([]*MXML, len(r.Pppppp))
	for i, c := range r.Pppppp {
		children["pppppp"][i] = c.ToMXML()
	}

	children["rf"] = make([]*MXML, len(r.Rf))
	for i, c := range r.Rf {
		children["rf"][i] = c.ToMXML()
	}

	children["rfz"] = make([]*MXML, len(r.Rfz))
	for i, c := range r.Rfz {
		children["rfz"][i] = c.ToMXML()
	}

	children["sf"] = make([]*MXML, len(r.Sf))
	for i, c := range r.Sf {
		children["sf"][i] = c.ToMXML()
	}

	children["sffz"] = make([]*MXML, len(r.Sffz))
	for i, c := range r.Sffz {
		children["sffz"][i] = c.ToMXML()
	}

	children["sfp"] = make([]*MXML, len(r.Sfp))
	for i, c := range r.Sfp {
		children["sfp"][i] = c.ToMXML()
	}

	children["sfpp"] = make([]*MXML, len(r.Sfpp))
	for i, c := range r.Sfpp {
		children["sfpp"][i] = c.ToMXML()
	}

	children["sfz"] = make([]*MXML, len(r.Sfz))
	for i, c := range r.Sfz {
		children["sfz"][i] = c.ToMXML()
	}

	children["sfzp"] = make([]*MXML, len(r.Sfzp))
	for i, c := range r.Sfzp {
		children["sfzp"][i] = c.ToMXML()
	}

	return newMXML("dynamics", strings.TrimSpace(r.IValue), attributes, children)
}
