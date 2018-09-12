package musicxml

import (
	"errors"
	"strconv"
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
	m.ppTimeSignatures()

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
	works, ok := m.MXML().Children()["work"]
	if !ok || len(works) == 0 {
		return errors.New("cannot get work title")
	}

	workTitles, ok := works[0].Children()["work-title"]
	if !ok || len(workTitles) == 0 {
		return errors.New("cannot get work title")
	}

	m.meta.workTitle = workTitles[0].IValue()

	return nil
}

func (m *MusicXML) ppCreators() error {
	identifications, ok := m.MXML().Children()["identification"]
	if !ok || len(identifications) == 0 {
		return errors.New("cannot get creators")
	}

	creators, ok := identifications[0].Children()["creator"]
	if !ok || len(creators) == 0 {
		return errors.New("cannot get creators")
	}

	m.meta.creators = make([]string, len(creators))
	for ix, c := range creators {
		m.meta.creators[ix] = c.IValue()
	}

	return nil
}

func (m *MusicXML) ppRights() error {
	identifications, ok := m.MXML().Children()["identification"]
	if !ok || len(identifications) == 0 {
		return errors.New("cannot get rights")
	}

	rights, ok := identifications[0].Children()["rights"]
	if !ok || len(rights) == 0 {
		return errors.New("cannot get rights")
	}

	m.meta.rights = make([]string, len(rights))
	for ix, r := range rights {
		m.meta.creators[ix] = r.IValue()
	}

	return nil
}

func (m *MusicXML) ppTimeSignatures() error {
	m.meta.timeSignatures = make([]*timeSignature, 0)

	parts, ok := m.MXML().Children()["part"]
	if !ok || len(parts) == 0 {
		return errors.New("cannot get time signatures")
	}

	measures, ok := parts[0].Children()["measure"]
	if !ok || len(measures) == 0 {
		return errors.New("cannot get time signatures")
	}

	measuresCount := len(measures)

	for ix, ms := range measures {
		attributes, ok := ms.Children()["attributes"]
		if !ok || len(attributes) == 0 {
			continue
		}

		timeSigs, ok := attributes[0].Children()["time"]
		if !ok || len(timeSigs) == 0 {
			continue
		}

		beats, ok := timeSigs[0].Children()["beats"]
		if !ok || len(beats) == 0 {
			return errors.New("cannot get time signatures")
		}

		beatType, ok := timeSigs[0].Children()["beat-type"]
		if !ok || len(beatType) == 0 {
			return errors.New("cannot get time signatures")
		}

		beatsInt, err := strconv.Atoi(beats[0].IValue())
		if err != nil {
			return errors.New("cannot get time signatures")
		}

		beatTypeInt, err := strconv.Atoi(beatType[0].IValue())
		if err != nil {
			return errors.New("cannot get time signatures")
		}

		m.meta.timeSignatures = append(m.meta.timeSignatures, &timeSignature{
			beats:            beatsInt,
			beatType:         beatTypeInt,
			lastMeasureIndex: measuresCount - 1,
		})

		if timeSigCount := len(m.meta.timeSignatures); timeSigCount > 1 {
			m.meta.timeSignatures[timeSigCount-2].lastMeasureIndex = ix
		}
	}

	return nil
}
