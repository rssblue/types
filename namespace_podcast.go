package types

import (
	"encoding/xml"
)

// PodcastNamespace is the Podcasting 2.0 namespace.
type PodcastNamespace string

func (ns *PodcastNamespace) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if ns != nil {
		if *ns == "" {
			return xml.Attr{}, nil
		}
		return xml.Attr{Name: xml.Name{Local: "xmlns:podcast"}, Value: string(*ns)}, nil
	}
	return xml.Attr{Name: xml.Name{Local: "xmlns:podcast"}, Value: "https://podcastindex.org/namespace/1.0"}, nil
}

// PodcastGUID is the global identifier for a podcast. Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#guid
type PodcastGUID string

func (podcastGUID PodcastGUID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		XMLName xml.Name `xml:"podcast:guid"`
		GUID    string   `xml:",chardata"`
	}{
		GUID: string(podcastGUID),
	}, start)
}

// Transcript denotes episode's transcript. Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#transcript
type Transcript struct {
	XMLName      xml.Name `xml:"podcast:transcript"`
	URL          string   `xml:"url,attr"`
	MimetypeName string   `xml:"type,attr"`
}

// Chapters denotes episode's chapters. Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#chapters
type Chapters struct {
	XMLName      xml.Name `xml:"podcast:chapters"`
	URL          string   `xml:"url,attr"`
	MimetypeName string   `xml:"type,attr"`
}

// Value enables to describe Value 4 Value payments. Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#value
type Value struct {
	XMLName    xml.Name `xml:"podcast:value"`
	Type       string   `xml:"type,attr"`
	Method     string   `xml:"method,attr"`
	Suggested  *float64 `xml:"suggested,attr"`
	Recipients []ValueRecipient
}

// Locked tells podcast hosting platforms whether they are allowed to import
// the feed. Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#locked
type Locked struct {
	XMLName  xml.Name `xml:"podcast:locked"`
	Owner    string
	IsLocked bool
}

func (l Locked) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	strBool := "no"
	if l.IsLocked {
		strBool = "yes"
	}
	return e.EncodeElement(struct {
		Owner    string `xml:"owner,attr"`
		IsLocked string `xml:",chardata"`
	}{
		Owner:    l.Owner,
		IsLocked: strBool,
	}, start)
}

// Location describes editorial focus of podcast's or episode's content. Read
// more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#location
type Location struct {
	XMLName  xml.Name `xml:"podcast:location"`
	Geo      *string  `xml:"geo,attr"`
	OSM      *string  `xml:"osm,attr"`
	Location string   `xml:",chardata"`
}

// Funding denotes donation/funding links. Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#value
type Funding struct {
	XMLName xml.Name `xml:"podcast:funding"`
	URL     string   `xml:"url,attr"`
	Caption string   `xml:",chardata"`
}

// ValueRecipient describes the recipient of Value 4 Value payments. Read more
// at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#value
type ValueRecipient struct {
	XMLName     xml.Name `xml:"podcast:valueRecipient"`
	Name        *string  `xml:"name,attr"`
	CustomKey   *string  `xml:"customKey,attr"`
	CustomValue *string  `xml:"customValue,attr"`
	Type        string   `xml:"type,attr"`
	Address     string   `xml:"address,attr"`
	Split       uint     `xml:"split,attr"`
	Fee         *bool    `xml:"bool,attr"`
}
