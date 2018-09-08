package mxml

// MXML stores MusicXML data as a tree.
type MXML struct {
	name       string
	iValue     string
	attributes map[string]string
	children   map[string][]*MXML
}

func newMXML(name string, iValue string, attributes map[string]string, children map[string][]*MXML) *MXML {
	return &MXML{
		name:       name,
		iValue:     iValue,
		attributes: attributes,
		children:   children,
	}
}

// Name returns the name of the MusicXML element.
func (m *MXML) Name() string {
	return m.name
}

// IValue returns the inner value of the MusicXML element.
func (m *MXML) IValue() string {
	return m.iValue
}

// Attributes returns the attributes of the MusicXML element.
func (m *MXML) Attributes() map[string]string {
	return m.attributes
}

// Children returns the children of the MusicXML element.
func (m *MXML) Children() map[string][]*MXML {
	return m.children
}
