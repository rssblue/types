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
	XMLName         xml.Name    `xml:"channel"`
	Title           string      `xml:"title"`
	Description     Description `xml:"description"`
	ITunesImage     ITunesImage
	Language        string `xml:"language"`
	ITunesCategory  ITunesCategory
	ITunesExplicit  bool    `xml:"itunes:explicit"`
	Author          string  `xml:"itunes:author"`
	Website         *string `xml:"link"`
	ITunesOwner     ITunesOwner
	ITunesType      string  `xml:"itunes:type"`
	Copyright       *string `xml:"copyright"`
	PodcastLocked   *PodcastLocked
	PodcastLocation *PodcastLocation
	PodcastFundings []PodcastFunding
	PodcastValue    *PodcastValue
	PodcastGUID     *PodcastGUID `xml:"podcast:guid"`
	PodcastMedium   string       `xml:"podcast:medium"`
	Items           []Item
}

// Item represents episode of a podcast.
type Item struct {
	XMLName             xml.Name `xml:"item"`
	Title               string   `xml:"title"`
	Enclosure           Enclosure
	GUID                string       `xml:"guid"`
	PubDate             Date         `xml:"pubDate"`
	Description         *Description `xml:"description"`
	ITunesDuration      *int64       `xml:"itunes:duration"`
	Link                *string      `xml:"link"`
	ITunesImage         *ITunesImage
	ITunesExplicit      bool   `xml:"itunes:explicit"`
	ITunesEpisodeNumber *int64 `xml:"itunes:episode"`
	ITunesSeasonNumber  *int64 `xml:"itunes:season"`
	ITunesType          string `xml:"itunes:episodeType"`
	PodcastTranscript   *PodcastTranscript
	PodcastChapters     *PodcastChapters
	PodcastLocation     *PodcastLocation
	PodcastValue        *PodcastValue
}

// Date is used to format the publish date of an episode.
type Date time.Time

func (pd Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	t := time.Time(pd)
	v := t.Format("Mon, 2 Jan 2006 15:04:05 GMT")
	return e.EncodeElement(v, start)
}
