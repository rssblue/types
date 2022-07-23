package types

import (
	"encoding/xml"
	"time"
)

// RSS is the root element of podcast's RSS feed, denoting the namespaces and
// the version of the protocol.
type RSS struct {
	XMLName             xml.Name             `xml:"rss"`
	Version             RSSVersion           `xml:",attr"`
	ITunesNamespace     *ITunesNamespace     `xml:",attr"`
	GooglePlayNamespace *GooglePlayNamespace `xml:",attr"`
	ContentNamespace    *ContentNamespace    `xml:"xmlns:content,attr"`
	PodcastNamespace    *PodcastNamespace    `xml:"xmlns:podcast,attr"`
	Channel             Channel
}

// Channel represents the podcast's feed.
type Channel struct {
	XMLName     xml.Name    `xml:"channel"`
	Title       string      `xml:"title"`
	Description Description `xml:"description"`
	Image       Image
	Language    string `xml:"language"`
	Category    Category
	IsExplicit  bool    `xml:"itunes:explicit"`
	Author      string  `xml:"itunes:author"`
	Website     *string `xml:"link"`
	Owner       Owner
	Type        string  `xml:"itunes:type"`
	Copyright   *string `xml:"copyright"`
	Locked      *Locked
	Location    *Location
	Fundings    []Funding
	Value       *Value
	GUID        *PodcastGUID `xml:"podcast:guid"`
	Medium      string       `xml:"podcast:medium"`
	Items       []Item
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

// Owner is used for owner's contact information.
type Owner struct {
	XMLName      xml.Name `xml:"itunes:owner"`
	Name         string   `xml:"itunes:name"`
	EmailAddress string   `xml:"itunes:email"`
}

// Description is used for the podcast's or episode's description.
type Description string

func (d Description) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		Description string `xml:",cdata"`
	}{
		Description: string(d),
	}, start)
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

// Item represents episode of a podcast.
type Item struct {
	XMLName       xml.Name `xml:"item"`
	Title         string   `xml:"title"`
	Enclosure     Enclosure
	GUID          string       `xml:"guid"`
	PubDate       Date         `xml:"pubDate"`
	Description   *Description `xml:"description"`
	Duration      *int64       `xml:"itunes:duration"`
	Link          *string      `xml:"link"`
	Image         *Image
	IsExplicit    bool   `xml:"itunes:explicit"`
	EpisodeNumber *int64 `xml:"itunes:episode"`
	SeasonNumber  *int64 `xml:"itunes:season"`
	Type          string `xml:"itunes:episodeType"`
	Transcript    *Transcript
	Chapters      *Chapters
	Location      *Location
	Value         *Value
}

// Date is used to format the publish date of an episode.
type Date time.Time

func (pd Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	t := time.Time(pd)
	v := t.Format("Mon, 2 Jan 2006 15:04:05 GMT")
	return e.EncodeElement(v, start)
}

// Enclosure is used to link to the episode's media file.
type Enclosure struct {
	XMLName      xml.Name `xml:"enclosure"`
	URL          string   `xml:"url,attr"`
	Length       int64    `xml:"length,attr"`
	MimetypeName string   `xml:"type,attr"`
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
