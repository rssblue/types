package types

import (
	"encoding/xml"
	"strconv"
	"time"
)

// NamespaceITunes is the iTunes namespace.
const NamespaceITunes string = "http://www.itunes.com/dtds/podcast-1.0.dtd"

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

// ITunesDuration denotesthe duration of an episode.
type ITunesDuration time.Duration

func (d ITunesDuration) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	numSeconds := int(time.Duration(d).Seconds())

	return e.EncodeElement(struct {
		Duration string `xml:",chardata"`
	}{
		Duration: strconv.Itoa(numSeconds),
	}, start)
}
