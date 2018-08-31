package musicxml

import (
	"encoding/xml"
	"strings"
)

// Credit is the structure for credit MusicXML element.
type Credit struct {
	XMLName xml.Name `xml:"credit"`

	IdAttr   string `xml:"id,attr,omitempty"`
	PageAttr string `xml:"page,attr,omitempty"`

	Bookmark     []Bookmark     `xml:"bookmark,omitempty"`
	CreditImage  []CreditImage  `xml:"credit-image,omitempty"`
	CreditSymbol []CreditSymbol `xml:"credit-symbol,omitempty"`
	CreditType   []CreditType   `xml:"credit-type,omitempty"`
	CreditWords  []CreditWords  `xml:"credit-words,omitempty"`
	Link         []Link         `xml:"link,omitempty"`

	IValue string `xml:",chardata"`
}

// ToMXML creates a MXML.
func (r *Credit) ToMXML() *MXML {
	attributes := make(map[string]string)
	children := make(map[string][]*MXML)

	attributes["id"] = r.IdAttr
	attributes["page"] = r.PageAttr

	children["bookmark"] = make([]*MXML, len(r.Bookmark))
	for i, c := range r.Bookmark {
		children["bookmark"][i] = c.ToMXML()
	}

	children["credit-image"] = make([]*MXML, len(r.CreditImage))
	for i, c := range r.CreditImage {
		children["credit-image"][i] = c.ToMXML()
	}

	children["credit-symbol"] = make([]*MXML, len(r.CreditSymbol))
	for i, c := range r.CreditSymbol {
		children["credit-symbol"][i] = c.ToMXML()
	}

	children["credit-type"] = make([]*MXML, len(r.CreditType))
	for i, c := range r.CreditType {
		children["credit-type"][i] = c.ToMXML()
	}

	children["credit-words"] = make([]*MXML, len(r.CreditWords))
	for i, c := range r.CreditWords {
		children["credit-words"][i] = c.ToMXML()
	}

	children["link"] = make([]*MXML, len(r.Link))
	for i, c := range r.Link {
		children["link"][i] = c.ToMXML()
	}

	return newMXML("credit", strings.TrimSpace(r.IValue), attributes, children)
}
