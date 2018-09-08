package musicxml

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/musica/musicxml/mxml"
)

const (
	uncompressedXML            = "application/vnd.recordare.musicxml+xml"
	errorExpectedScorePartwise = "expected element type <score-partwise>"
	errorExpectedScoreTimewise = "expected element type <score-timewise>"
)

type mxlContainer struct {
	Files []struct {
		FullPath  string `xml:"full-path,attr,omitempty"`
		MediaType string `xml:"media-type,attr,omitempty"`
	} `xml:"rootfiles>rootfile,omitempty"`
}

// ParseXMLBuffer parses MusicXML data from a io.Reader into a MXML struct.
func parseXMLBuffer(r io.Reader) (*mxml.MXML, error) {
	// The musicxml root may be either score-partwise or score-timewise.
	// So keep a backup of the buffer for a second try.
	var buf bytes.Buffer
	tr := io.TeeReader(r, &buf)

	mx, err := tryScorePartwise(xml.NewDecoder(tr))
	if err != nil {
		return nil, err
	} else if mx != nil {
		return mx, nil
	}

	// buf may not have all the data from the input io reader.
	// so use ioutil.readall to dump the input io reader into the secondary buffer.
	ioutil.ReadAll(tr)

	mx, err = tryScoreTimewise(xml.NewDecoder(&buf))
	if err != nil {
		return nil, err
	} else if mx != nil {
		return mx, nil
	}

	return nil, errorIncorrectRoot
}

// ParseXMLFile parses MusicXML data from a file into a MXML struct.
func parseXMLFile(filePath string) (*mxml.MXML, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return parseXMLBuffer(f)
}

// ParseMXLFile parses MusicXML data from a mxl file into a MXML struct.
func parseMXLFile(filePath string) (*mxml.MXML, error) {
	// TODO: confirm that MXL files contain only one uncompressed MusicXML file.
	z, err := zip.OpenReader(filePath)
	if err != nil {
		return nil, err
	}
	defer z.Close()

	// collect pointers to all files
	fileMap := make(map[string]*zip.File)
	for _, f := range z.File {
		fileMap[f.Name] = f
	}

	// parse the meta file
	mxlInfo := &mxlContainer{}
	metaFile, err := fileMap["META-INF/container.xml"].Open()
	if err != nil {
		return nil, err
	}
	defer metaFile.Close()
	err = xml.NewDecoder(metaFile).Decode(mxlInfo)
	if err != nil {
		return nil, err
	}

	for _, f := range mxlInfo.Files {
		// parse the file only if it is a uncompressed xml file.
		if f.MediaType == "" || f.MediaType == uncompressedXML {
			x, err := fileMap[f.FullPath].Open()
			if err != nil {
				return nil, err
			}
			defer x.Close()

			return parseXMLBuffer(x)
		}
	}

	return nil, errorEmptyMXL
}

func tryScorePartwise(d *xml.Decoder) (*mxml.MXML, error) {
	score := &mxml.ScorePartwise{}
	err := d.Decode(score)
	if err != nil {
		if strings.HasPrefix(err.Error(), errorExpectedScorePartwise) {
			return nil, nil
		}
		return nil, err
	}
	return score.ToMXML(), nil
}

func tryScoreTimewise(d *xml.Decoder) (*mxml.MXML, error) {
	score := &mxml.ScoreTimewise{}
	err := d.Decode(score)
	if err != nil {
		if strings.HasPrefix(err.Error(), errorExpectedScoreTimewise) {
			return nil, nil
		}
		return nil, err
	}
	return score.ToMXML(), nil
}
