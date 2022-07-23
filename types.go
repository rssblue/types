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
