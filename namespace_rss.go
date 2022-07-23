package types

import (
	"encoding/xml"
)

// RSSVersion denotes the RSS version.
type RSSVersion string

func (rssVersion RSSVersion) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if rssVersion != "" {
		return xml.Attr{Name: xml.Name{Local: "version"}, Value: string(rssVersion)}, nil
	}
	return xml.Attr{Name: xml.Name{Local: "version"}, Value: "2.0"}, nil
}

// Description is used for podcast's or episode's description.
type Description string

func (d Description) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		Description string `xml:",cdata"`
	}{
		Description: string(d),
	}, start)
}

// Enclosure is used to link to the episode's media file.
type Enclosure struct {
	XMLName      xml.Name `xml:"enclosure"`
	URL          string   `xml:"url,attr"`
	Length       int64    `xml:"length,attr"`
	MimetypeName string   `xml:"type,attr"`
}
