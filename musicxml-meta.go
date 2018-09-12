package musicxml

type musicXMLMeta struct {
	xmlRoot        string
	workTitle      string
	creators       []string
	rights         []string
	timeSignatures []*timeSignature
}

func newMusicXMLMeta() *musicXMLMeta {
	return &musicXMLMeta{}
}

// XMLRoot returns the name of the root MusicXML element. Right now it can return only score-partwise.
func (m *MusicXML) XMLRoot() string {
	return m.meta.xmlRoot
}

// WorkTitle returns the title of the score.
func (m *MusicXML) WorkTitle() string {
	return m.meta.workTitle
}

// Creators returns the creators of the score.
func (m *MusicXML) Creators() []string {
	return m.meta.creators
}

// Rights returns the copyright and other intellectual property notices of the score.
func (m *MusicXML) Rights() []string {
	return m.meta.rights
}
