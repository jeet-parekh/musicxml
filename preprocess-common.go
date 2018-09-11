package musicxml

func (m *MusicXML) preprocess() (*MusicXML, error) {
	// there would be a lot of preprocessing functions
	// don't need to return or panic in case one of them fails
	// figure out a way to handle them. log them??
	m.meta = newMusicXMLMeta()
	m.ppXMLRoot()
	m.ppSplit()
	return m, nil
}

// just to figure out an approach about how to preprocess the MusicXML.
func (m *MusicXML) ppXMLRoot() (*MusicXML, error) {
	switch mxRoot := m.musicXML.Name(); mxRoot {
	case "score-partwise":
		m.meta.xmlRoot = mxRoot
		return m, nil
	default:
		return m, errorIncorrectRoot
	}
}

// after all the common preprocessing has been done, do the individual preprocessing.
func (m *MusicXML) ppSplit() (*MusicXML, error) {
	switch m.meta.xmlRoot {
	case "score-partwise":
		return m.ppScorePartwise()
	default:
		return m, errorIncorrectRoot
	}
}
