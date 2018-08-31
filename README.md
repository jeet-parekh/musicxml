## MusicXML parser for Go.

---

- This package parses musicxml files into a simple tree-like structure. It provides the bare minimum functions needed to traverse the XML tree.

- To parse a musicxml file, use

  - `ParseXMLBytes(XMLData []byte)`.

  - `ParseXMLBuffer(XMLFileReader io.Reader)`.

  - `ParseXMLFile(XMLFilePath string)`.

  - `ParseMXLFile(CompressedXMLFilePath string)`
