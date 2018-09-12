package musicxml

import (
	"io"

	"github.com/musica/musicxml/mxml"
)

// MusicXML is a container for MusicXML data parsed into a MXML struct.
type MusicXML struct {
	musicXML *mxml.MXML
	meta     *musicXMLMeta
}

// NewMusicXML returns a new empty MusicXML struct.
func NewMusicXML() *MusicXML {
	return &MusicXML{}
}

// MXML returns the musicxml tree as MXML.
func (m *MusicXML) MXML() *mxml.MXML {
	return m.musicXML
}

// SetMXML sets a MXML to the MusicXML struct.
func (m *MusicXML) SetMXML(mx *mxml.MXML) (*MusicXML, error) {
	if root := mx.Name(); root == "score-partwise" {
		m.musicXML = mx
		return m.preprocess()
	}
	return m, errorIncorrectRoot
}

func (m *MusicXML) setMXML(mx *mxml.MXML, err error) (*MusicXML, error) {
	m.musicXML = mx
	if err != nil {
		return m, err
	}
	return m.preprocess()
}

// ParseXMLBuffer parses MusicXML data from a io.Reader into a MusicXML struct.
func ParseXMLBuffer(r io.Reader) (*MusicXML, error) {
	return NewMusicXML().setMXML(parseXMLBuffer(r))
}

// ParseXMLFile parses MusicXML data from a file into a MusicXML struct.
func ParseXMLFile(filePath string) (*MusicXML, error) {
	return NewMusicXML().setMXML(parseXMLFile(filePath))
}

// ParseMXLFile parses a MusicXML data from a mxl file into a MusicXML struct.
func ParseMXLFile(filePath string) (*MusicXML, error) {
	return NewMusicXML().setMXML(parseMXLFile(filePath))
}
