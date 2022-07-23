package types

import (
	"encoding/xml"
)

// ITunesNamespace is the iTunes namespace.
type ITunesNamespace string

func (ns *ITunesNamespace) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if ns != nil {
		if *ns == "" {
			return xml.Attr{}, nil
		}
		return xml.Attr{Name: xml.Name{Local: "xmlns:itunes"}, Value: string(*ns)}, nil
	}
	return xml.Attr{Name: xml.Name{Local: "xmlns:itunes"}, Value: "http://www.itunes.com/dtds/podcast-1.0.dtd"}, nil
}

// Owner is used for owner's contact information.
type Owner struct {
	XMLName      xml.Name `xml:"itunes:owner"`
	Name         string   `xml:"itunes:name"`
	EmailAddress string   `xml:"itunes:email"`
}

// Category denotes podcast's category information.
type Category struct {
	XMLName     xml.Name     `xml:"itunes:category"`
	Category    string       `xml:"text,attr"`
	Subcategory *Subcategory `xml:"itunes:category"`
}

// Subcategory is more granural; it is a subset of Category.
type Subcategory string

func (s Subcategory) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		Category string `xml:"text,attr"`
	}{
		Category: string(s),
	}, start)
}

// Image is podcast's or episode's artwork.
type Image struct {
	XMLName xml.Name `xml:"itunes:image"`
	URL     string   `xml:"href,attr"`
}
