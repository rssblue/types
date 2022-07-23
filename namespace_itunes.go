package types

import (
	"encoding/xml"
)

// NamespaceITunes is the iTunes namespace.
type NamespaceITunes string

func (ns *NamespaceITunes) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if ns != nil {
		if *ns == "" {
			return xml.Attr{}, nil
		}
		return xml.Attr{Name: xml.Name{Local: "xmlns:itunes"}, Value: string(*ns)}, nil
	}
	return xml.Attr{Name: xml.Name{Local: "xmlns:itunes"}, Value: "http://www.itunes.com/dtds/podcast-1.0.dtd"}, nil
}

// ITunesOwner is used for owner's contact information.
type ITunesOwner struct {
	XMLName xml.Name `xml:"itunes:owner"`
	Name    string   `xml:"itunes:name"`
	Email   string   `xml:"itunes:email"`
}

// ITunesCategory denotes podcast's category information.
type ITunesCategory struct {
	XMLName     xml.Name           `xml:"itunes:category"`
	Category    string             `xml:"text,attr"`
	Subcategory *ITunesSubcategory `xml:"itunes:category"`
}

// ITunesSubcategory is more granural; it is a subset of Category.
type ITunesSubcategory string

func (s ITunesSubcategory) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		Category string `xml:"text,attr"`
	}{
		Category: string(s),
	}, start)
}

// ITunesImage is podcast's or episode's artwork.
type ITunesImage struct {
	XMLName xml.Name `xml:"itunes:image"`
	URL     string   `xml:"href,attr"`
}
