package musicxml

import "errors"

var (
	errorPrefix = "musicxml: "

	errorIncorrectRoot = newError("MusicXML root element must be either score-partwise or score-timewise.")
	errorEmptyMusicXML = newError("no MusicXML data found.")
	errorEmptyMXL      = newError("no uncompressed MusicXML file found in the MXL file.")
)

func newError(errorMessage string) error {
	return errors.New(errorPrefix + errorMessage)
}
