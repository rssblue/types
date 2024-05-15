package types

import (
	"encoding/xml"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// NamespacePodcast is the Podcasting 2.0 namespace.
const NamespacePodcast string = "https://podcastindex.org/namespace/1.0"

// PodcastGUID is the global identifier for a podcast. Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#guid
type PodcastGUID string

// PodcastTranscript denotes episode's transcript. Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#transcript
type PodcastTranscript struct {
	XMLName  xml.Name `xml:"podcast:transcript"`
	URL      string   `xml:"url,attr"`
	Mimetype string   `xml:"type,attr"`
	Language *string  `xml:"language,attr"`
	Rel      *string  `xml:"rel,attr"`
}

// PodcastChapters denotes episode's chapters. Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#chapters
type PodcastChapters struct {
	XMLName  xml.Name `xml:"podcast:chapters"`
	URL      string   `xml:"url,attr"`
	Mimetype string   `xml:"type,attr"`
}

// PodcastValue enables to describe Value 4 Value payments. Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#value
type PodcastValue struct {
	XMLName         xml.Name `xml:"podcast:value"`
	Type            string   `xml:"type,attr"`
	Method          string   `xml:"method,attr"`
	Suggested       *float64 `xml:"suggested,attr,omitempty"`
	Recipients      []PodcastValueRecipient
	ValueTimeSplits []PodcastValueTimeSplit
}

// PodcastValueRecipient describes the recipient of Value 4 Value payments.
// Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#value
type PodcastValueRecipient struct {
	XMLName     xml.Name `xml:"podcast:valueRecipient"`
	Name        *string  `xml:"name,attr"`
	CustomKey   *string  `xml:"customKey,attr"`
	CustomValue *string  `xml:"customValue,attr"`
	Type        string   `xml:"type,attr"`
	Address     string   `xml:"address,attr"`
	Split       uint     `xml:"split,attr"`
	Fee         *bool    `xml:"bool,attr"`
}

// PodcastValueTimeSplit describes value splits that are valid for a certain period of time
// Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#value-time-split
type PodcastValueTimeSplit struct {
	XMLName          xml.Name         `xml:"podcast:valueTimeSplit"`
	StartTime        DurationInteger  `xml:"startTime,attr"`
	Duration         DurationInteger  `xml:"duration,attr"`
	RemoteStartTime  *DurationInteger `xml:"remoteStartTime,attr,omitempty"`
	RemotePercentage *uint            `xml:"remotePercentage,attr,omitempty"`
	Recipients       []PodcastValueRecipient
	RemoteItem       PodcastRemoteItem
}

// PodcastRemoteItem provides a way to "point" to another feed or item in it.
// Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#remote-item
type PodcastRemoteItem struct {
	XMLName  xml.Name       `xml:"podcast:remoteItem"`
	ItemGUID *string        `xml:"itemGuid,attr"`
	FeedGUID uuid.UUID      `xml:"feedGuid,attr"`
	FeedURL  *string        `xml:"feedUrl,attr"`
	Medium   *PodcastMedium `xml:"medium,attr"`
}

// PodcastLocked tells podcast hosting platforms whether they are allowed to import
// the feed. Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#locked
type PodcastLocked struct {
	XMLName  xml.Name `xml:"podcast:locked"`
	Owner    *string
	IsLocked bool
}

func (l PodcastLocked) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	strBool := "no"
	if l.IsLocked {
		strBool = "yes"
	}
	return e.EncodeElement(struct {
		Owner    *string `xml:"owner,attr"`
		IsLocked string  `xml:",chardata"`
	}{
		Owner:    l.Owner,
		IsLocked: strBool,
	}, start)
}

// PodcastLocation describes editorial focus of podcast's or episode's content.
// Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#location
type PodcastLocation struct {
	XMLName  xml.Name    `xml:"podcast:location"`
	Geo      *PodcastGeo `xml:",attr,omitempty"`
	OSM      *PodcastOSM `xml:",attr,omitempty"`
	Location string      `xml:",chardata"`
}

type PodcastGeo struct {
	Latitude    float64
	Longitude   float64
	Altitude    *float64
	Uncertainty *float64
}

// PodcastGeo is a geo URI, conforming to RFC 5870.
func (geo PodcastGeo) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	s := fmt.Sprintf("geo:%s,%s", removeTrailingZeros(geo.Latitude), removeTrailingZeros(geo.Longitude))
	if geo.Altitude != nil {
		s += fmt.Sprintf(",%s", removeTrailingZeros(*geo.Altitude))
	}
	if geo.Uncertainty != nil {
		s += fmt.Sprintf(";u=%s", removeTrailingZeros(*geo.Uncertainty))
	}

	return xml.Attr{Name: xml.Name{Local: "geo"}, Value: s}, nil
}

// PodcastOSM encodes OpenStreetMap location information.
type PodcastOSM struct {
	Type      rune
	FeatureID uint
	Revision  *uint
}

func (osm PodcastOSM) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	s := fmt.Sprintf("%c%d", osm.Type, osm.FeatureID)
	if osm.Revision != nil {
		s += fmt.Sprintf("#%d", *osm.Revision)
	}

	return xml.Attr{Name: xml.Name{Local: "osm"}, Value: s}, nil
}

func removeTrailingZeros(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

// PodcastFunding denotes donation/funding links. Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#value
type PodcastFunding struct {
	XMLName xml.Name `xml:"podcast:funding"`
	URL     string   `xml:"url,attr"`
	Caption string   `xml:",chardata"`
}

// PodcastSoundbite denotes soundbite associated with an episode. Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#soundbite
type PodcastSoundbite struct {
	XMLName   xml.Name `xml:"podcast:soundbite"`
	StartTime Duration `xml:"startTime,attr"`
	Duration  Duration `xml:"duration,attr"`
	Title     *string  `xml:",chardata"`
}

// Duration denotes timestamps and durations during a podcast episode.
type Duration time.Duration

const (
	Hour        = Duration(time.Hour)
	Minute      = Duration(time.Minute)
	Second      = Duration(time.Second)
	Millisecond = Duration(time.Millisecond)
	Microsecond = Duration(time.Microsecond)
	Nanosecond  = Duration(time.Nanosecond)
)

func (duration Duration) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	minutes := time.Duration(duration).Seconds()
	s := removeTrailingZeros(minutes)
	// // Add ".0" if does not exist.
	if !strings.Contains(s, ".") {
		s += ".0"
	}

	return xml.Attr{Name: xml.Name{Local: name.Local}, Value: s}, nil
}

// DurationInteger denotes timestamps and durations during a podcast episode, but which are converted to integer seconds.
type DurationInteger time.Duration

func (duration DurationInteger) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	seconds := time.Duration(duration).Seconds()
	seconds = math.Round(seconds)
	s := strconv.Itoa(int(seconds))

	return xml.Attr{Name: xml.Name{Local: name.Local}, Value: s}, nil
}

// PodcastPerson specifies a person of interest to the podcast. Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#person
type PodcastPerson struct {
	XMLName  xml.Name `xml:"podcast:person"`
	Name     string   `xml:",chardata"`
	Group    *string  `xml:"group,attr"`
	Role     *string  `xml:"role,attr"`
	URL      *string  `xml:"href,attr"`
	ImageURL *string  `xml:"img,attr"`
}

// PodcastSeason is used for identifying which episodes in a podcast are part
// of a particular "season". Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#season
type PodcastSeason struct {
	XMLName xml.Name `xml:"podcast:season"`
	Number  int      `xml:",chardata"`
	Name    *string  `xml:"name,attr"`
}

// PodcastEpisode exists largely for compatibility with PodcastSeason.
// Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#season
type PodcastEpisode struct {
	XMLName xml.Name `xml:"podcast:episode"`
	Number  float64  `xml:",chardata"`
	Display *string  `xml:"display,attr"`
}

// PodcastTrailer is used to define the location of an audio or video file to
// be used as a trailer for the entire podcast or a specific season. Read more
// at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#trailer
type PodcastTrailer struct {
	XMLName  xml.Name `xml:"podcast:trailer"`
	Title    string   `xml:",chardata"`
	PubDate  Date     `xml:"pubdate,attr"`
	URL      string   `xml:"url,attr"`
	Length   *int64   `xml:"length,attr"`
	Mimetype *string  `xml:"type,attr"`
	Season   *int     `xml:"season,attr"`
}

// PodcastMedium tells what the content contained within the feed is. Read more
// at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#medium
type PodcastMedium string

var (
	PodcastMediumPodcast    PodcastMedium = "podcast"
	PodcastMediumMusic      PodcastMedium = "music"
	PodcastMediumVideo      PodcastMedium = "video"
	PodcastMediumFilm       PodcastMedium = "film"
	PodcastMediumAudioBook  PodcastMedium = "audiobook"
	PodcastMediumNewsletter PodcastMedium = "newsletter"
	PodcastMediumBlog       PodcastMedium = "blog"

	PodcastMediumPublisher PodcastMedium = "publisher"

	PodcastMediumPodcastList    PodcastMedium = "podcastL"
	PodcastMediumMusicList      PodcastMedium = "musicL"
	PodcastMediumVideoList      PodcastMedium = "videoL"
	PodcastMediumFilmList       PodcastMedium = "filmL"
	PodcastMediumAudioBookList  PodcastMedium = "audiobookL"
	PodcastMediumNewsletterList PodcastMedium = "newsletterL"
	PodcastMediumBlogList       PodcastMedium = "blogL"
	PodcastMediumPublisherList  PodcastMedium = "publisherL"
	PodcastMediumMixedList      PodcastMedium = "mixed"
)

// PodcastTXT is intended for free-form text and is modeled after the DNS "TXT"
// record. Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#txt
type PodcastTXT struct {
	XMLName xml.Name `xml:"podcast:txt"`
	TXT     string   `xml:",chardata"`
	Purpose *string  `xml:"purpose,attr"`
}

// PodcastISRC is an experimental tag to store International Standard Recording
// Codes. Read more at https://isrc.ifpi.org
type PodcastISRC struct {
	XMLName xml.Name `xml:"podcast:isrc"`
	ISRC    string   `xml:",chardata"`
}

// PodcastPodping allows feed owners to signal to aggregators that the feed sends out Podping notifications when changes are made to it.
// Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#podping
type PodcastPodping struct {
	XMLName     xml.Name `xml:"podcast:podping"`
	UsesPodping *bool    `xml:"usesPodping,attr"`
}

// PodcastPublisher allows a podcast feed to link to it's "publisher feed" parent.
// Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#publisher
type PodcastPublisher struct {
	XMLName     xml.Name `xml:"podcast:publisher"`
	RemoteItems []PodcastRemoteItem
}

// PodcastAlternateEnclosure provides different versions of, or companion media to the main `<enclosure>` file.
// Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#alternate-enclosure
type PodcastAlternateEnclosure struct {
	XMLName      xml.Name `xml:"podcast:alternateEnclosure"`
	Mimetype     string   `xml:"type,attr"`
	Length       *int64   `xml:"length,attr"`
	Bitrate      *int64   `xml:"bitrate,attr"`
	Height       *int64   `xml:"height,attr"`
	LanguageCode *string  `xml:"lang,attr"`
	Title        *string  `xml:"title,attr"`
	Rel          *string  `xml:"rel,attr"`
	Default      *bool    `xml:"default,attr"`
	Sources      []PodcastSource
}

// PodcastSource defines a uri location for a `<podcast:alternateEnclosure>` media file.
// Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#source
type PodcastSource struct {
	XMLName     xml.Name `xml:"podcast:source"`
	URI         string   `xml:"uri,attr"`
	ContentType *string  `xml:"contentType,attr"`
}

type PodcastContentLink struct {
	XMLName xml.Name `xml:"podcast:contentLink"`
	Href    string   `xml:"href,attr"`
	Text    string   `xml:",chardata"`
}

type PodcastLiveStatus string

var (
	PodcastLiveStatusPending PodcastLiveStatus = "pending"
	PodcastLiveStatusLive    PodcastLiveStatus = "live"
	PodcastLiveStatusEnded   PodcastLiveStatus = "ended"
)

type PodcastLiveItem struct {
	XMLName xml.Name `xml:"podcast:liveItem"`

	Status    PodcastLiveStatus `xml:"status,attr"`
	StartTime time.Time         `xml:"start,attr"`
	EndTime   *time.Time        `xml:"end,attr,omitempty"`

	Description                *Description `xml:"description"`
	Enclosure                  *Enclosure
	GUID                       *GUID
	Link                       *string `xml:"link"`
	Title                      *string `xml:"title"`
	ContentEncoded             *ContentEncoded
	ITunesEpisodeNumber        *int64  `xml:"itunes:episode"`
	ITunesEpisodeType          *string `xml:"itunes:episodeType"`
	ITunesExplicit             *bool   `xml:"itunes:explicit"`
	ITunesImage                *ITunesImage
	ITunesSeasonNumber         *int64 `xml:"itunes:season"`
	PodcastAlternateEnclosures []PodcastAlternateEnclosure
	PodcastChat                *PodcastChat
	PodcastContentLinks        []PodcastContentLink
	PodcastEpisode             *PodcastEpisode
	PodcastISRC                *PodcastISRC
	PodcastLiveValue           *PodcastLiveValue
	PodcastLocation            *PodcastLocation
	PodcastPersons             []PodcastPerson
	PodcastSeason              *PodcastSeason
	PodcastSoundbites          []PodcastSoundbite
	PodcastTXTs                []PodcastTXT
	PodcastTranscripts         []PodcastTranscript
	PodcastValue               *PodcastValue
}

// PodcastLiveValue is an experimental tag to transmit updates during a livestream.
type PodcastLiveValue struct {
	XMLName  xml.Name `xml:"podcast:liveValue"`
	URI      string   `xml:"uri,attr"`
	Protocol string   `xml:"protocol,attr"`
}

// PodcastChat is an experimental tag to enable chat during a livestream.
type PodcastChat struct {
	XMLName   xml.Name `xml:"podcast:chat"`
	Server    string   `xml:"server,attr"`
	Protocol  string   `xml:"protocol,attr"`
	AccountID *string  `xml:"accountId,attr"`
	Space     *string  `xml:"space,attr"`
	EmbedURL  *string  `xml:"embedUrl,attr"`
}

// PodcastSingleItem denotes whether the feed contains a single item or multiple items.
// It's a proposal described at https://github.com/Podcastindex-org/podcast-namespace/discussions/578
type PodcastSingleItem struct {
	XMLName xml.Name `xml:"podcast:singleItem"`
	Value   bool     `xml:",chardata"`
}
