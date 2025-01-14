package types

import (
	"encoding/xml"
	"fmt"
	"time"
)

// RSS is the root element of podcast's RSS feed, denoting the namespaces and
// the version of the protocol.
type RSS struct {
	XMLName             xml.Name   `xml:"rss"`
	Version             RSSVersion `xml:",attr"`
	NamespaceAtom       NSBool     `xml:",attr"`
	NamespaceContent    NSBool     `xml:",attr"`
	NamespaceGooglePlay NSBool     `xml:",attr"`
	NamespaceITunes     NSBool     `xml:",attr"`
	NamespacePodcast    NSBool     `xml:",attr"`
	NamespacePSC        NSBool     `xml:",attr"`
	Channel             Channel
}

type NSBool bool

func (isPresent *NSBool) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if isPresent == nil {
		return xml.Attr{}, nil
	}

	if !*isPresent {
		return xml.Attr{}, nil
	}

	switch name.Local {
	case "NamespaceAtom":
		return xml.Attr{Name: xml.Name{Local: "xmlns:atom"}, Value: NamespaceAtom}, nil
	case "NamespaceContent":
		return xml.Attr{Name: xml.Name{Local: "xmlns:content"}, Value: NamespaceContent}, nil
	case "NamespaceGooglePlay":
		return xml.Attr{Name: xml.Name{Local: "xmlns:googleplay"}, Value: NamespaceGooglePlay}, nil
	case "NamespaceITunes":
		return xml.Attr{Name: xml.Name{Local: "xmlns:itunes"}, Value: NamespaceITunes}, nil
	case "NamespacePodcast":
		return xml.Attr{Name: xml.Name{Local: "xmlns:podcast"}, Value: NamespacePodcast}, nil
	case "NamespacePSC":
		return xml.Attr{Name: xml.Name{Local: "xmlns:psc"}, Value: NamespacePSC}, nil
	default:
		return xml.Attr{}, fmt.Errorf("unrecognised attribute name \"%s\"", name.Local)
	}
}

// Channel represents the podcast's feed.
type Channel struct {
	XMLName            xml.Name     `xml:"channel"`
	Copyright          *string      `xml:"copyright"`
	Description        *Description `xml:"description"`
	Generator          *string      `xml:"generator"`
	Language           *string      `xml:"language"`
	LastBuildDate      *Date        `xml:"lastBuildDate"`
	Link               *string      `xml:"link"`
	Title              *string      `xml:"title"`
	AtomLink           *AtomLink    `xml:"atom:link"`
	ContentEncoded     *ContentEncoded
	ITunesAuthor       *string `xml:"itunes:author"`
	ITunesCategories   []ITunesCategory
	ITunesExplicit     *bool `xml:"itunes:explicit"`
	ITunesImage        *ITunesImage
	ITunesNewFeedURL   *string `xml:"itunes:new-feed-url"`
	ITunesOwner        *ITunesOwner
	ITunesType         *string `xml:"itunes:type"`
	PodcastFundings    []PodcastFunding
	PodcastGUID        *PodcastGUID `xml:"podcast:guid"`
	PodcastLicense     *PodcastLicense
	PodcastLocation    *PodcastLocation
	PodcastLocked      *PodcastLocked
	PodcastMedium      *PodcastMedium `xml:"podcast:medium"`
	PodcastPersons     []PodcastPerson
	PodcastPodping     *PodcastPodping
	PodcastPublisher   *PodcastPublisher
	PodcastRemoteItems []PodcastRemoteItem
	PodcastSingleItem  *PodcastSingleItem
	PodcastTXTs        []PodcastTXT
	PodcastTrailers    []PodcastTrailer
	PodcastValue       *PodcastValue
	PodcastLiveItems   []PodcastLiveItem
	Items              []Item
}

// Item represents episode of a podcast.
type Item struct {
	XMLName                    xml.Name     `xml:"item"`
	Description                *Description `xml:"description"`
	Enclosure                  *Enclosure
	GUID                       *GUID
	Link                       *string `xml:"link"`
	PubDate                    *Date   `xml:"pubDate"`
	Title                      *string `xml:"title"`
	ContentEncoded             *ContentEncoded
	ITunesDuration             *ITunesDuration `xml:"itunes:duration"`
	ITunesEpisodeNumber        *int64          `xml:"itunes:episode"`
	ITunesEpisodeType          *string         `xml:"itunes:episodeType"`
	ITunesExplicit             *bool           `xml:"itunes:explicit"`
	ITunesImage                *ITunesImage
	ITunesSeasonNumber         *int64 `xml:"itunes:season"`
	PodcastAlternateEnclosures []PodcastAlternateEnclosure
	PodcastChapters            *PodcastChapters
	PodcastEpisode             *PodcastEpisode
	PodcastISRC                *PodcastISRC
	PodcastLicense             *PodcastLicense
	PodcastLocation            *PodcastLocation
	PodcastPersons             []PodcastPerson
	PodcastSeason              *PodcastSeason
	PodcastSoundbites          []PodcastSoundbite
	PodcastTXTs                []PodcastTXT
	PodcastTranscripts         []PodcastTranscript
	PodcastValue               *PodcastValue
	PSCChapters                *PSCChapters
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
