package rss

import (
	"encoding/xml"
	"time"
)

type RSS struct {
	XMLName           xml.Name `xml:"rss"`
	Version           string   `xml:"version,attr"`
	XMLNSITunes       string   `xml:"xmlns:itunes,attr"`
	XMLNSGooglePlay   string   `xml:"xmlns:googleplay,attr"`
	XMLNSContent      string   `xml:"xmlns:content,attr"`
	XMLNSPodcasting20 string   `xml:"xmlns:podcast,attr"`
	Channel           Channel
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description Description
	Image       ChannelImage
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
	GUID        PodcastGUID
	Medium      string `xml:"podcast:medium"`
	Items       []Item
}

type PodcastGUID struct {
	XMLName xml.Name `xml:"podcast:guid,omitempty"`
	GUID    string   `xml:",chardata"`
}

type Owner struct {
	XMLName      xml.Name `xml:"itunes:owner"`
	Name         string   `xml:"itunes:name"`
	EmailAddress string   `xml:"itunes:email"`
}

type Description struct {
	XMLName     xml.Name `xml:"description"`
	Description string   `xml:",cdata"`
}

type Category struct {
	XMLName     xml.Name `xml:"itunes:category"`
	Category    string   `xml:"text,attr"`
	Subcategory *Subcategory
}

type Subcategory struct {
	XMLName     xml.Name `xml:"itunes:category"`
	Subcategory string   `xml:"text,attr"`
}

type ChannelImage struct {
	XMLName xml.Name `xml:"itunes:image"`
	Href    string   `xml:"href,attr"`
}

type Item struct {
	XMLName       xml.Name `xml:"item"`
	Title         string   `xml:"title"`
	Enclosure     Enclosure
	GUID          string  `xml:"guid"`
	PubDate       PubDate `xml:"pubDate"`
	Description   *Description
	Duration      *int64  `xml:"itunes:duration"`
	Link          *string `xml:"link"`
	Image         *ItemImage
	IsExplicit    bool   `xml:"itunes:explicit"`
	EpisodeNumber *int64 `xml:"itunes:episode"`
	SeasonNumber  *int64 `xml:"itunes:season"`
	Type          string `xml:"itunes:episodeType"`
	Transcript    *Transcript
	Chapters      *Chapters
	Location      *Location
	Value         *Value
}

type PubDate time.Time

const pubDateFormat = "Mon, 2 Jan 2006 15:04:05 GMT"

func (pd PubDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	t := time.Time(pd)
	v := t.Format(pubDateFormat)
	return e.EncodeElement(v, start)
}

func (pd *PubDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}
	t, err := time.Parse(pubDateFormat, s)
	if err != nil {
		return err
	}
	*pd = PubDate(t)
	return nil
}

type Enclosure struct {
	XMLName      xml.Name `xml:"enclosure"`
	URL          string   `xml:"url,attr"`
	Length       int64    `xml:"length,attr"`
	MimetypeName string   `xml:"type,attr"`
}

// See <https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md>
type Transcript struct {
	XMLName      xml.Name `xml:"podcast:transcript,omitempty"`
	URL          string   `xml:"url,attr"`
	MimetypeName string   `xml:"type,attr"`
}

// See <https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md>
type Chapters struct {
	XMLName      xml.Name `xml:"podcast:chapters,omitempty"`
	URL          string   `xml:"url,attr"`
	MimetypeName string   `xml:"type,attr"`
}

// See <https://github.com/Podcastindex-org/podcast-namespace/blob/main/value/value.md>
type Value struct {
	XMLName    xml.Name `xml:"podcast:value,omitempty"`
	Type       string   `xml:"type,attr"`
	Method     string   `xml:"method,attr"`
	Suggested  *float64 `xml:"suggested,attr"`
	Recipients []ValueRecipient
}

// See <https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md>
type Locked struct {
	XMLName  xml.Name `xml:"podcast:locked,omitempty"`
	Owner    string
	IsLocked bool
}

type locked struct {
	Owner    string `xml:"owner,attr"`
	IsLocked string `xml:",chardata"`
}

func (l Locked) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	strBool := "no"
	if l.IsLocked {
		strBool = "yes"
	}
	return e.EncodeElement(locked{Owner: l.Owner, IsLocked: strBool}, start)
}

func (l *Locked) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var isLockedStruct locked
	err := d.DecodeElement(&isLockedStruct, &start)
	if err != nil {
		return err
	}
	if isLockedStruct.IsLocked == "yes" {
		l.IsLocked = true
	}
	l.Owner = isLockedStruct.Owner
	return nil
}

// See <https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md>
type Location struct {
	XMLName  xml.Name `xml:"podcast:location,omitempty"`
	Geo      *string  `xml:"geo,attr"`
	OSM      *string  `xml:"osm,attr"`
	Location string   `xml:",chardata"`
}

// See <https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md>
type Funding struct {
	XMLName xml.Name `xml:"podcast:funding,omitempty"`
	URL     string   `xml:"url,attr"`
	Caption string   `xml:",chardata"`
}

// See <https://github.com/Podcastindex-org/podcast-namespace/blob/main/value/value.md>
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

type ItemImage struct {
	XMLName xml.Name `xml:"itunes:image,omitempty"`
	Href    string   `xml:"href,attr"`
}