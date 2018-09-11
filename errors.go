package musicxml

import "errors"

var (
	errorPrefix = "musicxml: "

	errorScoreTimewise = newError("score-timewise is not yet supported.")

	errorIncorrectRoot = newError("MusicXML root element must be score-partwise.")
	errorEmptyMusicXML = newError("no MusicXML data found.")
	errorEmptyMXL      = newError("no uncompressed MusicXML file found in the MXL file.")
)

func newError(errorMessage string) error {
	return errors.New(errorPrefix + errorMessage)
}
