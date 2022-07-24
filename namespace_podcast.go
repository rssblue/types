package types

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// NamespacePodcast is the Podcasting 2.0 namespace.
type NamespacePodcast string

func (ns *NamespacePodcast) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
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
	XMLName    xml.Name `xml:"podcast:value"`
	Type       string   `xml:"type,attr"`
	Method     string   `xml:"method,attr"`
	Suggested  *float64 `xml:"suggested,attr,omitempty"`
	Recipients []PodcastValueRecipient
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

// PodcastLocked tells podcast hosting platforms whether they are allowed to import
// the feed. Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#locked
type PodcastLocked struct {
	XMLName  xml.Name `xml:"podcast:locked"`
	Owner    string
	IsLocked bool
}

func (l PodcastLocked) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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
	Title     *string  `xml:",innerxml"`
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

// PodcastPerson specifies a person of interest to the podcast. Read more at
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#person
type PodcastPerson struct {
	XMLName  xml.Name `xml:"podcast:person"`
	Name     string   `xml:",innerxml"`
	Group    *string  `xml:"group,attr"`
	Role     *string  `xml:"role,attr"`
	URL      *string  `xml:"href,attr"`
	ImageURL *string  `xml:"img,attr"`
}
