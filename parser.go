package musicxml

import (
	"archive/zip"
	"encoding/xml"
	"io"
	"os"

	"github.com/musica/musicxml/mxml"
)

const (
	uncompressedXML = "application/vnd.recordare.musicxml+xml"
)

type mxlContainer struct {
	Files []struct {
		FullPath  string `xml:"full-path,attr,omitempty"`
		MediaType string `xml:"media-type,attr,omitempty"`
	} `xml:"rootfiles>rootfile,omitempty"`
}

// ParseXMLBuffer parses MusicXML data from a io.Reader into a MXML struct.
func parseXMLBuffer(r io.Reader) (*mxml.MXML, error) {
	score := &mxml.ScorePartwise{}
	err := xml.NewDecoder(r).Decode(score)
	if err != nil {
		return nil, err
	}
	return score.ToMXML(), nil
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
