package types

import (
	"encoding/xml"
)

// NamespaceGooglePlay is the Google Play namespace.
type NamespaceGooglePlay string

func (ns *NamespaceGooglePlay) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if ns != nil {
		if *ns == "" {
			return xml.Attr{}, nil
		}
		return xml.Attr{Name: xml.Name{Local: "xmlns:googleplay"}, Value: string(*ns)}, nil
	}
	return xml.Attr{Name: xml.Name{Local: "xmlns:googleplay"}, Value: "http://www.google.com/schemas/play-podcasts/1.0"}, nil
}
