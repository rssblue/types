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
type Description struct {
	XMLName     xml.Name `xml:"description"`
	Description string
	IsCDATA     bool
}

func (d Description) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if d.IsCDATA {
		return e.EncodeElement(struct {
			Description string `xml:",cdata"`
		}{
			Description: d.Description,
		}, start)
	}
	return e.EncodeElement(struct {
		Description string `xml:",innerxml"`
	}{
		Description: d.Description,
	}, start)
}

// Enclosure is used to link to the episode's media file.
type Enclosure struct {
	XMLName      xml.Name `xml:"enclosure"`
	URL          string   `xml:"url,attr"`
	Length       int64    `xml:"length,attr"`
	MimetypeName string   `xml:"type,attr"`
}
