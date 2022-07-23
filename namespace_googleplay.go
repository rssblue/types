package types

import (
	"encoding/xml"
)

// GooglePlayNamespace is the Google Play namespace.
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
