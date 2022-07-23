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
	Copyright       *string     `xml:"copyright"`
	Description     Description `xml:"description"`
	Language        string      `xml:"language"`
	Link            *string     `xml:"link"`
	Title           string      `xml:"title"`
	ITunesAuthor    string      `xml:"itunes:author"`
	ITunesCategory  ITunesCategory
	ITunesExplicit  bool `xml:"itunes:explicit"`
	ITunesImage     ITunesImage
	ITunesOwner     ITunesOwner
	ITunesType      string `xml:"itunes:type"`
	PodcastFundings []PodcastFunding
	PodcastGUID     *PodcastGUID `xml:"podcast:guid"`
	PodcastLocation *PodcastLocation
	PodcastLocked   *PodcastLocked
	PodcastMedium   string `xml:"podcast:medium"`
	PodcastValue    *PodcastValue
	Items           []Item
}

// Item represents episode of a podcast.
type Item struct {
	XMLName             xml.Name     `xml:"item"`
	Description         *Description `xml:"description"`
	Enclosure           Enclosure
	GUID                string  `xml:"guid"`
	Link                *string `xml:"link"`
	PubDate             Date    `xml:"pubDate"`
	Title               string  `xml:"title"`
	ITunesDuration      *int64  `xml:"itunes:duration"`
	ITunesEpisodeNumber *int64  `xml:"itunes:episode"`
	ITunesEpisodeType   string  `xml:"itunes:episodeType"`
	ITunesExplicit      bool    `xml:"itunes:explicit"`
	ITunesImage         *ITunesImage
	ITunesSeasonNumber  *int64 `xml:"itunes:season"`
	PodcastChapters     *PodcastChapters
	PodcastLocation     *PodcastLocation
	PodcastTranscript   *PodcastTranscript
	PodcastValue        *PodcastValue
}

// Date is used to format the publish date of an episode.
type Date time.Time

func (pd Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	t := time.Time(pd)
	v := t.Format("Mon, 2 Jan 2006 15:04:05 GMT")
	return e.EncodeElement(v, start)
}
