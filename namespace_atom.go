package types

import (
	"encoding/xml"
)

// NamespaceAtom is the Atom namespace.
const NamespaceAtom string = "http://www.w3.org/2005/Atom"

// AtomLink defines a reference from an entry or feed to a Web resource.
type AtomLink struct {
	XMLName xml.Name `xml:"atom:link"`
	Href    string   `xml:"href,attr"`
	Rel     *string  `xml:"rel,attr,omitempty"`
	Type    *string  `xml:"type,attr,omitempty"`
}
