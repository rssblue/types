package types_test

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/rssblue/types"
)

func TestMarshal(t *testing.T) {
	tests := []struct {
		unmarshalled types.RSS
		marshalled   string
	}{
		{
			unmarshalled: types.RSS{
				Channel: types.Channel{
					Title:       "Bookworm Podcast",
					Description: types.Description("<strong>Description</strong>"),
					Image: types.Image{
						Href: "https://rssblue.com/@bookworm-podcast/cover-art.png",
					},
					Language: "en",
					Category: types.Category{
						Category: "Society & Culture",
						Subcategory: &types.Subcategory{
							Subcategory: "Documentary",
						},
					},
					IsExplicit: true,
					Author:     "Jane Doe",
					Website:    pointer("https://example.com"),
					Owner: types.Owner{
						Name:         "Jane Doe",
						EmailAddress: "jane@example.com",
					},
					Type:      "episodic",
					Copyright: pointer("© RSS Blue"),
					Locked: &types.Locked{
						Owner:    "jane@example.com",
						IsLocked: false,
					},
					Fundings: []types.Funding{
						{
							URL:     "https://example.com/donate",
							Caption: "Support “Bookworm Podcast”",
						},
					},
					Value: &types.Value{
						Type:   "lightning",
						Method: "keysend",
						Recipients: []types.ValueRecipient{
							{
								Name:    pointer("Co-Host #1"),
								Type:    "node",
								Address: "02d5c1bf8b940dc9cadca86d1b0a3c37fbe39cee4c7e839e33bef9174531d27f52",
								Split:   50,
							},
							{
								Name:    pointer("Co-Host #2"),
								Type:    "node",
								Address: "032f4ffbbafffbe51726ad3c164a3d0d37ec27bc67b29a159b0f49ae8ac21b8508",
								Split:   40,
							},
							{
								Name:    pointer("Producer"),
								Type:    "node",
								Address: "03ae9f91a0cb8ff43840e3c322c4c61f019d8c1c3cea15a25cfc425ac605e61a4a",
								Split:   10,
							},
						},
					},
					GUID: types.PodcastGUID{
						GUID: "cda647ce-56b8-5d7c-9448-ba1993ab46b7",
					},
					Medium: "podcast",
					Items: []types.Item{
						{
							Title: "Hello Again",
							Enclosure: types.Enclosure{
								URL:          "https://rssblue.com/@bookworm-podcast/hello-again/hello-again.mp3",
								MimetypeName: "audio/mpeg",
								Length:       2048,
							},
							GUID:    "hello-again",
							PubDate: types.PubDate(time.Date(2021, time.July, 10, 9, 3, 59, 0, time.UTC)),
							Image: &types.Image{
								Href: "https://rssblue.com/@bookworm-podcast/hello-again/cover-art.png",
							},
							Type: "full",
							Transcript: &types.Transcript{
								URL:          "https://rssblue.com/@bookworm-podcast/hello-again/transcript.vtt",
								MimetypeName: "text/vtt",
							},
							Value: &types.Value{
								Type:   "lightning",
								Method: "keysend",
								Recipients: []types.ValueRecipient{
									{
										Name:    pointer("Host"),
										Type:    "node",
										Address: "02d5c1bf8b940dc9cadca86d1b0a3c37fbe39cee4c7e839e33bef9174531d27f52",
										Split:   90,
									},
									{
										Name:    pointer("Producer"),
										Type:    "node",
										Address: "03ae9f91a0cb8ff43840e3c322c4c61f019d8c1c3cea15a25cfc425ac605e61a4a",
										Split:   10,
									},
								},
							},
						},
						{
							Title: "Hello World",
							Enclosure: types.Enclosure{
								URL:          "https://rssblue.com/@bookworm-podcast/hello-world/hello-world.mp3",
								MimetypeName: "audio/mpeg",
								Length:       1024,
							},
							GUID:        "hello-world",
							PubDate:     types.PubDate(time.Date(2021, time.July, 8, 15, 20, 10, 0, time.UTC)),
							Description: pointer(types.Description("This is my <em>first</em> episode!")),
							IsExplicit:  true,
							Type:        "full",
							Transcript: &types.Transcript{
								URL:          "https://rssblue.com/@bookworm-podcast/hello-world/transcript.srt",
								MimetypeName: "application/x-subrip",
							},
							Chapters: &types.Chapters{
								URL:          "https://rssblue.com/@bookworm-podcast/hello-world/chapters.json",
								MimetypeName: "application/json+chapters",
							},
							Location: &types.Location{
								Geo:      pointer("geo:39.7837304,-100.445882;u=3900000"),
								OSM:      pointer("R148838"),
								Location: "Gitmo Nation",
							},
						},
					},
				},
			},
			marshalled: `<rss version="2.0" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:googleplay="http://www.google.com/schemas/play-podcasts/1.0" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:podcast="https://podcastindex.org/namespace/1.0">
  <channel>
    <title>Bookworm Podcast</title>
    <description><![CDATA[<strong>Description</strong>]]></description>
    <itunes:image href="https://rssblue.com/@bookworm-podcast/cover-art.png"></itunes:image>
    <language>en</language>
    <itunes:category text="Society &amp; Culture">
      <itunes:category text="Documentary"></itunes:category>
    </itunes:category>
    <itunes:explicit>true</itunes:explicit>
    <itunes:author>Jane Doe</itunes:author>
    <link>https://example.com</link>
    <itunes:owner>
      <itunes:name>Jane Doe</itunes:name>
      <itunes:email>jane@example.com</itunes:email>
    </itunes:owner>
    <itunes:type>episodic</itunes:type>
    <copyright>© RSS Blue</copyright>
    <podcast:locked owner="jane@example.com">no</podcast:locked>
    <podcast:funding url="https://example.com/donate">Support “Bookworm Podcast”</podcast:funding>
    <podcast:value type="lightning" method="keysend">
      <podcast:valueRecipient name="Co-Host #1" type="node" address="02d5c1bf8b940dc9cadca86d1b0a3c37fbe39cee4c7e839e33bef9174531d27f52" split="50"></podcast:valueRecipient>
      <podcast:valueRecipient name="Co-Host #2" type="node" address="032f4ffbbafffbe51726ad3c164a3d0d37ec27bc67b29a159b0f49ae8ac21b8508" split="40"></podcast:valueRecipient>
      <podcast:valueRecipient name="Producer" type="node" address="03ae9f91a0cb8ff43840e3c322c4c61f019d8c1c3cea15a25cfc425ac605e61a4a" split="10"></podcast:valueRecipient>
    </podcast:value>
    <podcast:guid>cda647ce-56b8-5d7c-9448-ba1993ab46b7</podcast:guid>
    <podcast:medium>podcast</podcast:medium>
    <item>
      <title>Hello Again</title>
      <enclosure url="https://rssblue.com/@bookworm-podcast/hello-again/hello-again.mp3" length="2048" type="audio/mpeg"></enclosure>
      <guid>hello-again</guid>
      <pubDate>Sat, 10 Jul 2021 09:03:59 GMT</pubDate>
      <itunes:image href="https://rssblue.com/@bookworm-podcast/hello-again/cover-art.png"></itunes:image>
      <itunes:explicit>false</itunes:explicit>
      <itunes:episodeType>full</itunes:episodeType>
      <podcast:transcript url="https://rssblue.com/@bookworm-podcast/hello-again/transcript.vtt" type="text/vtt"></podcast:transcript>
      <podcast:value type="lightning" method="keysend">
        <podcast:valueRecipient name="Host" type="node" address="02d5c1bf8b940dc9cadca86d1b0a3c37fbe39cee4c7e839e33bef9174531d27f52" split="90"></podcast:valueRecipient>
        <podcast:valueRecipient name="Producer" type="node" address="03ae9f91a0cb8ff43840e3c322c4c61f019d8c1c3cea15a25cfc425ac605e61a4a" split="10"></podcast:valueRecipient>
      </podcast:value>
    </item>
    <item>
      <title>Hello World</title>
      <enclosure url="https://rssblue.com/@bookworm-podcast/hello-world/hello-world.mp3" length="1024" type="audio/mpeg"></enclosure>
      <guid>hello-world</guid>
      <pubDate>Thu, 8 Jul 2021 15:20:10 GMT</pubDate>
      <description><![CDATA[This is my <em>first</em> episode!]]></description>
      <itunes:explicit>true</itunes:explicit>
      <itunes:episodeType>full</itunes:episodeType>
      <podcast:transcript url="https://rssblue.com/@bookworm-podcast/hello-world/transcript.srt" type="application/x-subrip"></podcast:transcript>
      <podcast:chapters url="https://rssblue.com/@bookworm-podcast/hello-world/chapters.json" type="application/json+chapters"></podcast:chapters>
      <podcast:location geo="geo:39.7837304,-100.445882;u=3900000" osm="R148838">Gitmo Nation</podcast:location>
    </item>
  </channel>
</rss>`,
		},
		{
			unmarshalled: types.RSS{
				ContentNamespace: pointer(types.ContentNamespace("")), // Should remove the namespace.
				Channel: types.Channel{
					Title:       "World Explorer Podcast",
					Description: types.Description("Very interesting podcast."),
					Image: types.Image{
						Href: "https://rssblue.com/@world-explorer-podcast/cover-art.jpg",
					},
					Language: "fr",
					Category: types.Category{
						Category: "Fiction",
					},
					Author: "John Doe",
					Owner: types.Owner{
						Name:         "John Doe",
						EmailAddress: "john@example.com",
					},
					Type: "serial",
					Location: &types.Location{
						Geo:      pointer("geo:30.2672,97.7431"),
						OSM:      pointer("R113314"),
						Location: "Austin, TX",
					},
					GUID: types.PodcastGUID{
						GUID: "96b952d9-06b2-5489-a3f3-d371473121fa",
					},
					Medium: "music",
				},
			},
			marshalled: `<rss version="2.0" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:googleplay="http://www.google.com/schemas/play-podcasts/1.0" xmlns:podcast="https://podcastindex.org/namespace/1.0">
  <channel>
    <title>World Explorer Podcast</title>
    <description><![CDATA[Very interesting podcast.]]></description>
    <itunes:image href="https://rssblue.com/@world-explorer-podcast/cover-art.jpg"></itunes:image>
    <language>fr</language>
    <itunes:category text="Fiction"></itunes:category>
    <itunes:explicit>false</itunes:explicit>
    <itunes:author>John Doe</itunes:author>
    <itunes:owner>
      <itunes:name>John Doe</itunes:name>
      <itunes:email>john@example.com</itunes:email>
    </itunes:owner>
    <itunes:type>serial</itunes:type>
    <podcast:location geo="geo:30.2672,97.7431" osm="R113314">Austin, TX</podcast:location>
    <podcast:guid>96b952d9-06b2-5489-a3f3-d371473121fa</podcast:guid>
    <podcast:medium>music</podcast:medium>
  </channel>
</rss>`,
		},
	}

	for i, test := range tests {
		// Marshalling
		marshalled, err := xml.MarshalIndent(&test.unmarshalled, "", "  ")
		if err != nil {
			t.Errorf("%d: unexpected error: %v", i, err)
		}
		diff := cmp.Diff(test.marshalled, string(marshalled))
		if diff != "" {
			t.Errorf("%d: mismatch (-want +got):\n%s", i, diff)
		}
	}
}

func pointer[T any](v T) *T {
	return &v
}
