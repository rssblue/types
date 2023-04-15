package types

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"time"
)

// NamespacePSC is the namespace for Podlove Simple Chapters.
const NamespacePSC string = "http://podlove.org/simple-chapters"

// PSCChapters is the root element for Podlove Simple Chapters.
type PSCChapters struct {
	XMLName  xml.Name `xml:"psc:chapters"`
	Version  string   `xml:"version,attr"`
	Chapters []PSCChapter
}

// PSCChapter is a single chapter in Podlove Simple Chapters.
type PSCChapter struct {
	Start time.Duration
	Title string
	Href  *url.URL
	Image *url.URL
}

func (encoded PSCChapter) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	// Do default except for attributes, which we will marshal ourselves.
	start.Name.Local = "psc:chapter"

	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "start"},
		Value: formatChapterStart(encoded.Start),
	})
	start.Attr = append(start.Attr, xml.Attr{
		Name:  xml.Name{Local: "title"},
		Value: encoded.Title,
	})
	if encoded.Href != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "href"},
			Value: encoded.Href.String(),
		})
	}
	if encoded.Image != nil {
		start.Attr = append(start.Attr, xml.Attr{
			Name:  xml.Name{Local: "image"},
			Value: encoded.Image.String(),
		})
	}
	return e.EncodeElement(struct{}{}, start)
}

func formatChapterStart(start time.Duration) string {
	hours := int(start.Hours())
	minutes := int(start.Minutes()) - hours*60
	seconds := int(start.Seconds()) - minutes*60 - hours*3600
	milliseconds := int(start.Milliseconds()) - seconds*1000 - minutes*60000 - hours*3600000

	str := fmt.Sprintf("%02d:%02d", minutes, seconds)
	if hours > 0 {
		str = fmt.Sprintf("%02d:%s", hours, str)
	}
	if milliseconds > 0 {
		str = fmt.Sprintf("%s.%03d", str, milliseconds)
	}

	return str
}
