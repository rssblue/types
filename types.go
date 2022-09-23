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
	NamespaceContent    *NamespaceContent    `xml:"xmlns:content,attr"`
	NamespaceGooglePlay *NamespaceGooglePlay `xml:",attr"`
	NamespaceITunes     *NamespaceITunes     `xml:",attr"`
	NamespacePodcast    *NamespacePodcast    `xml:"xmlns:podcast,attr"`
	Channel             Channel
}

// Channel represents the podcast's feed.
type Channel struct {
	XMLName          xml.Name    `xml:"channel"`
	Copyright        *string     `xml:"copyright"`
	Description      Description `xml:"description"`
	Generator        *string     `xml:"generator"`
	Language         string      `xml:"language"`
	Link             *string     `xml:"link"`
	Title            string      `xml:"title"`
	ContentEncoded   *ContentEncoded
	ITunesAuthor     string `xml:"itunes:author"`
	ITunesCategories []ITunesCategory
	ITunesExplicit   bool `xml:"itunes:explicit"`
	ITunesImage      ITunesImage
	ITunesOwner      ITunesOwner
	ITunesType       string `xml:"itunes:type"`
	PodcastFundings  []PodcastFunding
	PodcastGUID      *PodcastGUID `xml:"podcast:guid"`
	PodcastLocation  *PodcastLocation
	PodcastLocked    *PodcastLocked
	PodcastMedium    string `xml:"podcast:medium"`
	PodcastPersons   []PodcastPerson
	PodcastTrailers  []PodcastTrailer
	PodcastValue     *PodcastValue
	Items            []Item
}

// Item represents episode of a podcast.
type Item struct {
	XMLName             xml.Name     `xml:"item"`
	Description         *Description `xml:"description"`
	Enclosure           Enclosure
	GUID                GUID
	Link                *string `xml:"link"`
	PubDate             Date    `xml:"pubDate"`
	Title               string  `xml:"title"`
	ContentEncoded      *ContentEncoded
	ITunesDuration      *int64 `xml:"itunes:duration"`
	ITunesEpisodeNumber *int64 `xml:"itunes:episode"`
	ITunesEpisodeType   string `xml:"itunes:episodeType"`
	ITunesExplicit      *bool  `xml:"itunes:explicit"`
	ITunesImage         *ITunesImage
	ITunesSeasonNumber  *int64 `xml:"itunes:season"`
	PodcastChapters     *PodcastChapters
	PodcastEpisode      *PodcastEpisode
	PodcastLocation     *PodcastLocation
	PodcastPersons      []PodcastPerson
	PodcastSeason       *PodcastSeason
	PodcastSoundbites   []PodcastSoundbite
	PodcastTranscripts  []PodcastTranscript
	PodcastValue        *PodcastValue
}

// Date is used to format the publish date of an episode.
type Date time.Time

func (pd Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	t := time.Time(pd)
	v := t.Format("Mon, 02 Jan 2006 15:04:05 GMT")
	return e.EncodeElement(v, start)
}

func (pd Date) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	t := time.Time(pd)
	v := t.Format("Mon, 02 Jan 2006 15:04:05 GMT")
	return xml.Attr{Name: name, Value: v}, nil
}
