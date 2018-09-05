## MusicXML parser for Go.

---

- This package parses MusicXML files into a simple tree-like structure. It provides the bare minimum functions needed to traverse the XML tree.

- To parse a MusicXML file, use

  - `ParseXMLBuffer(XMLFileReader io.Reader)`.

  - `ParseXMLFile(XMLFilePath string)`.

  - `ParseMXLFile(MXLFilePath string)`
