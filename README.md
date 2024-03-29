# Types

This package provides a number of Go struct types with field tags for XML marshalling.
There are standard RSS 2.0, iTunes and many of the [Podcasting 2.0](https://github.com/Podcastindex-org/podcast-namespace) tags available.

## Install

There is no stable release yet, and backwards-incompatible changes may still be introduced.
But if you want to try it, you can simply do
```
go get github.com/rssblue/types
```

## Example

### Code

```go
package main

import (
  "encoding/xml"
  "fmt"
  "time"

  "github.com/rssblue/types"
)

func main() {
  rss := types.RSS{
    NamespaceITunes:  true,
    NamespacePodcast: true,
    Channel: types.Channel{
      Title: pointer("Bookworm Podcast"),
      Description: &types.Description{
        Description: "Podcast about <em>books</em>.",
        IsCDATA:     true,
      },
      Language:     pointer("en"),
      ITunesAuthor: pointer("John"),
      ITunesOwner: &types.ITunesOwner{
        Name:  "John",
        Email: "john@example.com",
      },
      ITunesImage: &types.ITunesImage{
        URL: "https://example.com/cover-art.png",
      },
      ITunesCategories: []types.ITunesCategory{
        {
          Category: "Arts",
        },
      },
      ITunesType:  pointer("episodic"),
      PodcastGUID: pointer(types.PodcastGUID("cda647ce-56b8-5d7c-9448-ba1993ab46b7")),
      Items: []types.Item{
        {
          Title: pointer("Book Review: Moby-Dick"),
          Enclosure: &types.Enclosure{
            URL:      "https://example.com/moby-dick.mp3",
            Length:   4096,
            Mimetype: "audio/mpeg",
          },
          GUID:              &types.GUID{GUID: "https://example.com/moby-dick"},
          ITunesEpisodeType: pointer("full"),
          PubDate:           pointer(types.Date(time.Date(2022, time.July, 23, 10, 30, 0, 0, time.UTC))),
          PodcastLocation: &types.PodcastLocation{
            OSM: &types.PodcastOSM{
              Type:      'R',
              FeatureID: 2396248,
            },
          },
        },
      },
    },
  }

  output, err := xml.MarshalIndent(&rss, "", "  ")
  if err != nil {
    panic(err)
  }

  fmt.Printf("%s%s\n", xml.Header, output)

}

func pointer[T any](v T) *T {
  return &v
}
```

### Output

```xml
<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:podcast="https://podcastindex.org/namespace/1.0">
  <channel>
    <description><![CDATA[Podcast about <em>books</em>.]]></description>
    <language>en</language>
    <title>Bookworm Podcast</title>
    <itunes:author>John</itunes:author>
    <itunes:category text="Arts"></itunes:category>
    <itunes:image href="https://example.com/cover-art.png"></itunes:image>
    <itunes:owner>
      <itunes:name>John</itunes:name>
      <itunes:email>john@example.com</itunes:email>
    </itunes:owner>
    <itunes:type>episodic</itunes:type>
    <podcast:guid>cda647ce-56b8-5d7c-9448-ba1993ab46b7</podcast:guid>
    <item>
      <enclosure url="https://example.com/moby-dick.mp3" length="4096" type="audio/mpeg"></enclosure>
      <guid>https://example.com/moby-dick</guid>
      <pubDate>Sat, 23 Jul 2022 10:30:00 GMT</pubDate>
      <title>Book Review: Moby-Dick</title>
      <itunes:episodeType>full</itunes:episodeType>
      <podcast:location osm="R2396248"></podcast:location>
    </item>
  </channel>
</rss>
```
