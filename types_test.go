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
					Title: "Bookworm Podcast",
					Description: types.Description{
						Description: "<strong>Description</strong>",
						IsCDATA:     true,
					},
					ITunesImage: types.ITunesImage{
						URL: "https://rssblue.com/@bookworm-podcast/cover-art.png",
					},
					Language: "en",
					ITunesCategories: []types.ITunesCategory{
						{
							Category:    "Society & Culture",
							Subcategory: pointer(types.ITunesSubcategory("Documentary")),
						},
					},
					ITunesExplicit: true,
					ITunesAuthor:   "Jane Doe",
					Link:           pointer("https://example.com"),
					ITunesOwner: types.ITunesOwner{
						Name:  "Jane Doe",
						Email: "jane@example.com",
					},
					ITunesType: "episodic",
					Copyright:  pointer("© RSS Blue"),
					PodcastLocked: &types.PodcastLocked{
						Owner:    "jane@example.com",
						IsLocked: false,
					},
					PodcastFundings: []types.PodcastFunding{
						{
							URL:     "https://example.com/donate",
							Caption: "Support “Bookworm Podcast”",
						},
					},
					PodcastValue: &types.PodcastValue{
						Type:   "lightning",
						Method: "keysend",
						Recipients: []types.PodcastValueRecipient{
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
					PodcastGUID:   pointer(types.PodcastGUID("cda647ce-56b8-5d7c-9448-ba1993ab46b7")),
					PodcastMedium: "podcast",
					PodcastPersons: []types.PodcastPerson{
						{
							Name:     "John Smith",
							URL:      pointer("https://example.com/johnsmith/blog"),
							ImageURL: pointer("http://example.com/images/johnsmith.jpg"),
						},
					},
					Items: []types.Item{
						{
							Title: "Hello Again",
							Enclosure: types.Enclosure{
								URL:      "https://rssblue.com/@bookworm-podcast/hello-again/hello-again.mp3",
								Mimetype: "audio/mpeg",
								Length:   2048,
							},
							GUID:    "hello-again",
							PubDate: types.Date(time.Date(2021, time.July, 10, 9, 3, 59, 0, time.UTC)),
							ITunesImage: &types.ITunesImage{
								URL: "https://rssblue.com/@bookworm-podcast/hello-again/cover-art.png",
							},
							ITunesEpisodeType: "full",
							PodcastTranscripts: []types.PodcastTranscript{
								{
									URL:      "https://rssblue.com/@bookworm-podcast/hello-again/transcript.vtt",
									Mimetype: "text/vtt",
								},
							},
							PodcastValue: &types.PodcastValue{
								Type:   "lightning",
								Method: "keysend",
								Recipients: []types.PodcastValueRecipient{
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
							PodcastSoundbites: []types.PodcastSoundbite{
								{
									StartTime: types.Minute + 13*types.Second,
									Duration:  types.Minute,
								},
								{
									StartTime: 20*types.Minute + 34*types.Second + 500*types.Millisecond,
									Duration:  42*types.Second + 250*types.Millisecond,
									Title:     pointer("Why the Podcast Namespace Matters"),
								},
							},
							PodcastPersons: []types.PodcastPerson{
								{
									Name:     "Jane Doe",
									Role:     pointer("guest"),
									URL:      pointer("https://www.imdb.com/name/nm0427852888/"),
									ImageURL: pointer("http://example.com/images/janedoe.jpg"),
								},
								{
									Name:     "Alice Brown",
									Role:     pointer("guest"),
									URL:      pointer("https://www.wikipedia/alicebrown"),
									ImageURL: pointer("http://example.com/images/alicebrown.jpg"),
								},
							},
							PodcastSeason: &types.PodcastSeason{
								Number: 5,
							},
						},
						{
							Title: "Hello World",
							Enclosure: types.Enclosure{
								URL:      "https://rssblue.com/@bookworm-podcast/hello-world/hello-world.mp3",
								Mimetype: "audio/mpeg",
								Length:   1024,
							},
							GUID:    "hello-world",
							PubDate: types.Date(time.Date(2021, time.July, 8, 15, 20, 10, 0, time.UTC)),
							Description: &types.Description{
								Description: "This is my <em>first</em> episode!",
								IsCDATA:     true,
							},
							ITunesExplicit:    true,
							ITunesEpisodeType: "full",
							PodcastTranscripts: []types.PodcastTranscript{
								{
									URL:      "https://rssblue.com/@bookworm-podcast/hello-world/transcript.srt",
									Mimetype: "application/x-subrip",
								},
							},
							PodcastChapters: &types.PodcastChapters{
								URL:      "https://rssblue.com/@bookworm-podcast/hello-world/chapters.json",
								Mimetype: "application/json+chapters",
							},
							PodcastLocation: &types.PodcastLocation{
								Geo: &types.PodcastGeo{
									Latitude:    39.7837304,
									Longitude:   -100.445882,
									Uncertainty: pointer(3900000.0),
								},
								OSM: &types.PodcastOSM{
									Type:      'R',
									FeatureID: 148838,
								},
								Location: "Gitmo Nation",
							},
							PodcastPersons: []types.PodcastPerson{
								{
									Name:     "Alice Brown",
									Role:     pointer("guest"),
									URL:      pointer("https://www.wikipedia/alicebrown"),
									ImageURL: pointer("http://example.com/images/alicebrown.jpg"),
									Group:    pointer("writing"),
								},
								{
									Name:  "Becky Smith",
									Role:  pointer("Cover Art Designer"),
									URL:   pointer("https://example.com/artist/beckysmith"),
									Group: pointer("visuals"),
								},
							},
							PodcastSeason: &types.PodcastSeason{
								Number: 3,
								Name:   pointer("Race for the Whitehouse 2020"),
							},
						},
					},
				},
			},
			marshalled: `<rss version="2.0" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:googleplay="http://www.google.com/schemas/play-podcasts/1.0" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:podcast="https://podcastindex.org/namespace/1.0">
  <channel>
    <copyright>© RSS Blue</copyright>
    <description><![CDATA[<strong>Description</strong>]]></description>
    <language>en</language>
    <link>https://example.com</link>
    <title>Bookworm Podcast</title>
    <itunes:author>Jane Doe</itunes:author>
    <itunes:category text="Society &amp; Culture">
      <itunes:category text="Documentary"></itunes:category>
    </itunes:category>
    <itunes:explicit>true</itunes:explicit>
    <itunes:image href="https://rssblue.com/@bookworm-podcast/cover-art.png"></itunes:image>
    <itunes:owner>
      <itunes:name>Jane Doe</itunes:name>
      <itunes:email>jane@example.com</itunes:email>
    </itunes:owner>
    <itunes:type>episodic</itunes:type>
    <podcast:funding url="https://example.com/donate">Support “Bookworm Podcast”</podcast:funding>
    <podcast:guid>cda647ce-56b8-5d7c-9448-ba1993ab46b7</podcast:guid>
    <podcast:locked owner="jane@example.com">no</podcast:locked>
    <podcast:medium>podcast</podcast:medium>
    <podcast:person href="https://example.com/johnsmith/blog" img="http://example.com/images/johnsmith.jpg">John Smith</podcast:person>
    <podcast:value type="lightning" method="keysend">
      <podcast:valueRecipient name="Co-Host #1" type="node" address="02d5c1bf8b940dc9cadca86d1b0a3c37fbe39cee4c7e839e33bef9174531d27f52" split="50"></podcast:valueRecipient>
      <podcast:valueRecipient name="Co-Host #2" type="node" address="032f4ffbbafffbe51726ad3c164a3d0d37ec27bc67b29a159b0f49ae8ac21b8508" split="40"></podcast:valueRecipient>
      <podcast:valueRecipient name="Producer" type="node" address="03ae9f91a0cb8ff43840e3c322c4c61f019d8c1c3cea15a25cfc425ac605e61a4a" split="10"></podcast:valueRecipient>
    </podcast:value>
    <item>
      <enclosure url="https://rssblue.com/@bookworm-podcast/hello-again/hello-again.mp3" length="2048" type="audio/mpeg"></enclosure>
      <guid>hello-again</guid>
      <pubDate>Sat, 10 Jul 2021 09:03:59 GMT</pubDate>
      <title>Hello Again</title>
      <itunes:episodeType>full</itunes:episodeType>
      <itunes:explicit>false</itunes:explicit>
      <itunes:image href="https://rssblue.com/@bookworm-podcast/hello-again/cover-art.png"></itunes:image>
      <podcast:person role="guest" href="https://www.imdb.com/name/nm0427852888/" img="http://example.com/images/janedoe.jpg">Jane Doe</podcast:person>
      <podcast:person role="guest" href="https://www.wikipedia/alicebrown" img="http://example.com/images/alicebrown.jpg">Alice Brown</podcast:person>
      <podcast:season>5</podcast:season>
      <podcast:soundbite startTime="73.0" duration="60.0"></podcast:soundbite>
      <podcast:soundbite startTime="1234.5" duration="42.25">Why the Podcast Namespace Matters</podcast:soundbite>
      <podcast:transcript url="https://rssblue.com/@bookworm-podcast/hello-again/transcript.vtt" type="text/vtt"></podcast:transcript>
      <podcast:value type="lightning" method="keysend">
        <podcast:valueRecipient name="Host" type="node" address="02d5c1bf8b940dc9cadca86d1b0a3c37fbe39cee4c7e839e33bef9174531d27f52" split="90"></podcast:valueRecipient>
        <podcast:valueRecipient name="Producer" type="node" address="03ae9f91a0cb8ff43840e3c322c4c61f019d8c1c3cea15a25cfc425ac605e61a4a" split="10"></podcast:valueRecipient>
      </podcast:value>
    </item>
    <item>
      <description><![CDATA[This is my <em>first</em> episode!]]></description>
      <enclosure url="https://rssblue.com/@bookworm-podcast/hello-world/hello-world.mp3" length="1024" type="audio/mpeg"></enclosure>
      <guid>hello-world</guid>
      <pubDate>Thu, 8 Jul 2021 15:20:10 GMT</pubDate>
      <title>Hello World</title>
      <itunes:episodeType>full</itunes:episodeType>
      <itunes:explicit>true</itunes:explicit>
      <podcast:chapters url="https://rssblue.com/@bookworm-podcast/hello-world/chapters.json" type="application/json+chapters"></podcast:chapters>
      <podcast:location geo="geo:39.7837304,-100.445882;u=3900000" osm="R148838">Gitmo Nation</podcast:location>
      <podcast:person group="writing" role="guest" href="https://www.wikipedia/alicebrown" img="http://example.com/images/alicebrown.jpg">Alice Brown</podcast:person>
      <podcast:person group="visuals" role="Cover Art Designer" href="https://example.com/artist/beckysmith">Becky Smith</podcast:person>
      <podcast:season name="Race for the Whitehouse 2020">3</podcast:season>
      <podcast:transcript url="https://rssblue.com/@bookworm-podcast/hello-world/transcript.srt" type="application/x-subrip"></podcast:transcript>
    </item>
  </channel>
</rss>`,
		},
		{
			unmarshalled: types.RSS{
				NamespaceContent: pointer(types.NamespaceContent("")), // Should remove the namespace.
				Channel: types.Channel{
					Title: "World Explorer Podcast",
					Description: types.Description{
						Description: "Very interesting podcast.",
					},
					ITunesImage: types.ITunesImage{
						URL: "https://rssblue.com/@world-explorer-podcast/cover-art.jpg",
					},
					Language: "fr",
					ITunesCategories: []types.ITunesCategory{
						{
							Category: "Fiction",
						},
						{
							Category:    "Society & Culture",
							Subcategory: pointer(types.ITunesSubcategory("Documentary")),
						},
					},
					ITunesAuthor: "John Doe",
					ITunesOwner: types.ITunesOwner{
						Name:  "John Doe",
						Email: "john@example.com",
					},
					ITunesType: "serial",
					PodcastLocation: &types.PodcastLocation{
						OSM: &types.PodcastOSM{
							Type:      'R',
							FeatureID: 113314,
						},
						Location: "Austin, TX",
					},
					PodcastGUID:   pointer(types.PodcastGUID("96b952d9-06b2-5489-a3f3-d371473121fa")),
					PodcastMedium: "music",
				},
			},
			marshalled: `<rss version="2.0" xmlns:googleplay="http://www.google.com/schemas/play-podcasts/1.0" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd" xmlns:podcast="https://podcastindex.org/namespace/1.0">
  <channel>
    <description>Very interesting podcast.</description>
    <language>fr</language>
    <title>World Explorer Podcast</title>
    <itunes:author>John Doe</itunes:author>
    <itunes:category text="Fiction"></itunes:category>
    <itunes:category text="Society &amp; Culture">
      <itunes:category text="Documentary"></itunes:category>
    </itunes:category>
    <itunes:explicit>false</itunes:explicit>
    <itunes:image href="https://rssblue.com/@world-explorer-podcast/cover-art.jpg"></itunes:image>
    <itunes:owner>
      <itunes:name>John Doe</itunes:name>
      <itunes:email>john@example.com</itunes:email>
    </itunes:owner>
    <itunes:type>serial</itunes:type>
    <podcast:guid>96b952d9-06b2-5489-a3f3-d371473121fa</podcast:guid>
    <podcast:location osm="R113314">Austin, TX</podcast:location>
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
