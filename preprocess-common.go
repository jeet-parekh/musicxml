package musicxml

import (
	"errors"
	"fmt"

	"github.com/musica/musicxml/mxml"
)

func (m *MusicXML) preprocess() (*MusicXML, error) {
	// there would be a lot of preprocessing functions
	// don't need to return or panic in case one of them fails
	// figure out a way to handle them. log them??
	m.meta = newMusicXMLMeta()
	if err := m.ppXMLRoot(); err != nil {
		return m, err
	}

	m.ppWorkTitle()
	m.ppCreators()
	m.ppRights()

	return m, nil
}

// just to figure out an approach about how to preprocess the MusicXML.
func (m *MusicXML) ppXMLRoot() error {
	switch mxRoot := m.musicXML.Name(); mxRoot {
	case "score-partwise":
		m.meta.xmlRoot = mxRoot
		return nil
	default:
		return errorIncorrectRoot
	}
}

func (m *MusicXML) ppWorkTitle() error {
	var n []*mxml.MXML
	var ok bool

	n, ok = m.MXML().Children()["work"]
	if !ok || len(n) == 0 {
		return errors.New("cannot get work title")
	}

	n, ok = n[0].Children()["work-title"]
	if !ok || len(n) == 0 {
		return errors.New("cannot get work title")
	}

	fmt.Println(n[0])

	m.meta.workTitle = n[0].IValue()

	return nil
}

func (m *MusicXML) ppCreators() error {
	var n []*mxml.MXML
	var ok bool

	n, ok = m.MXML().Children()["identification"]
	if !ok || len(n) == 0 {
		return errors.New("cannot get creators")
	}

	n, ok = n[0].Children()["creator"]
	if !ok || len(n) == 0 {
		return errors.New("cannot get creators")
	}

	m.meta.creators = make([]string, len(n))
	for ix, v := range n {
		m.meta.creators[ix] = v.IValue()
	}

	return nil
}

func (m *MusicXML) ppRights() error {
	var n []*mxml.MXML
	var ok bool

	n, ok = m.MXML().Children()["identification"]
	if !ok || len(n) == 0 {
		return errors.New("cannot get creators")
	}

	n, ok = n[0].Children()["rights"]
	if !ok || len(n) == 0 {
		return errors.New("cannot get creators")
	}

	m.meta.creators = make([]string, len(n))
	for ix, v := range n {
		m.meta.creators[ix] = v.IValue()
	}

	return nil
}
