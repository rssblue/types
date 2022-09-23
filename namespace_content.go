package types

import (
	"encoding/xml"
)

// NamespaceContent is the namespace for RSS format's content module.
type NamespaceContent string

func (ns *NamespaceContent) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if ns != nil {
		if *ns == "" {
			return xml.Attr{}, nil
		}
		return xml.Attr{Name: xml.Name{Local: "xmlns:content"}, Value: string(*ns)}, nil
	}
	return xml.Attr{Name: xml.Name{Local: "xmlns:content"}, Value: "http://purl.org/rss/1.0/modules/content/"}, nil
}

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
		Encoded string `xml:",innerxml"`
	}{
		Encoded: encoded.Encoded,
	}, start)
}
