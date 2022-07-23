package types

import "encoding/xml"

type ITunesNamespace string

func (ns *ITunesNamespace) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if ns != nil {
		if *ns == "" {
			return xml.Attr{}, nil
		}
		return xml.Attr{Name: xml.Name{Local: "xmlns:itunes"}, Value: string(*ns)}, nil
	}
	return xml.Attr{Name: xml.Name{Local: "xmlns:itunes"}, Value: "http://www.itunes.com/dtds/podcast-1.0.dtd"}, nil
}

type GooglePlayNamespace string

func (ns *GooglePlayNamespace) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if ns != nil {
		if *ns == "" {
			return xml.Attr{}, nil
		}
		return xml.Attr{Name: xml.Name{Local: "xmlns:googleplay"}, Value: string(*ns)}, nil
	}
	return xml.Attr{Name: xml.Name{Local: "xmlns:googleplay"}, Value: "http://www.google.com/schemas/play-podcasts/1.0"}, nil
}

type PodcastNamespace string

func (ns *PodcastNamespace) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if ns != nil {
		if *ns == "" {
			return xml.Attr{}, nil
		}
		return xml.Attr{Name: xml.Name{Local: "xmlns:podcast"}, Value: string(*ns)}, nil
	}
	return xml.Attr{Name: xml.Name{Local: "xmlns:podcast"}, Value: "https://podcastindex.org/namespace/1.0"}, nil
}

type ContentNamespace string

func (ns *ContentNamespace) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if ns != nil {
		if *ns == "" {
			return xml.Attr{}, nil
		}
		return xml.Attr{Name: xml.Name{Local: "xmlns:content"}, Value: string(*ns)}, nil
	}
	return xml.Attr{Name: xml.Name{Local: "xmlns:content"}, Value: "http://purl.org/rss/1.0/modules/content/"}, nil
}

type RSSVersion string

func (rssVersion RSSVersion) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if rssVersion != "" {
		return xml.Attr{Name: xml.Name{Local: "version"}, Value: string(rssVersion)}, nil
	}
	return xml.Attr{Name: xml.Name{Local: "version"}, Value: "2.0"}, nil
}
