package musicxml

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"os"
	"strings"
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

// ParseXMLBuffer parses musicxml data from a io.Reader into a MXML struct.
func ParseXMLBuffer(r io.Reader) (*MXML, error) {
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

	mx, err = tryScoreTimewise(xml.NewDecoder(&buf))
	if err != nil {
		return nil, err
	} else if mx != nil {
		return mx, nil
	}

	return nil, errors.New("mxml: musicxml root element must be score-partwise or score-timewise")
}

// ParseXMLFile parses musicxml data from a file into a MXML struct.
func ParseXMLFile(filePath string) (*MXML, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseXMLBuffer(f)
}

// ParseMXLFile parses musicxml data from a mxl file into a MXML struct.
func ParseMXLFile(filePath string) (*MXML, error) {
	// TODO: confirm that MXL files contain only one uncompressed musicxml file.
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

			return ParseXMLBuffer(x)
		}
	}

	return nil, errors.New("mxml: no uncompressed musicxml file found in the mxl file")
}

func tryScorePartwise(d *xml.Decoder) (*MXML, error) {
	score := &ScorePartwise{}
	err := d.Decode(score)
	if err != nil {
		if strings.HasSuffix(err.Error(), errorExpectedScorePartwise) {
			return nil, nil
		}
		return nil, err
	}
	return score.ToMXML(), nil
}

func tryScoreTimewise(d *xml.Decoder) (*MXML, error) {
	score := &ScoreTimewise{}
	err := d.Decode(score)
	if err != nil {
		if strings.HasSuffix(err.Error(), errorExpectedScoreTimewise) {
			return nil, nil
		}
		return nil, err
	}
	return score.ToMXML(), nil
}
