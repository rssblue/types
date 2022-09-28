package types

import (
	"encoding/xml"
)

// NamespaceContent is the namespace for RSS format's content module.
const NamespaceContent string = "http://purl.org/rss/1.0/modules/content/"

// ContentEncoded is used for podcast's or episode's description.
type ContentEncoded struct {
	XMLName xml.Name `xml:"content:encoded"`
	Encoded string
	IsCDATA bool
}

func (encoded ContentEncoded) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if encoded.IsCDATA {
		return e.EncodeElement(struct {
			Encoded string `xml:",cdata"`
		}{
			Encoded: encoded.Encoded,
		}, start)
	}
	return e.EncodeElement(struct {
		Encoded string `xml:",chardata"`
	}{
		Encoded: encoded.Encoded,
	}, start)
}
