package vimego

import "testing"

func TestMetadata(t *testing.T) {
	video, _ := NewVideo("https://vimeo.com/206152466")
	metadata, err := video.Metadata()
	if err != nil {
		t.Fatal(err)
	}

	if metadata.Title != "Crystal Castles - Kept" {
		t.Error("metadata.Title doesn't match")
	}
	if metadata.Duration != 243 {
		t.Error("metadata.Duration doesn't match")
	}
}

func TestFormats(t *testing.T) {
	video, _ := NewVideo("https://vimeo.com/206152466")
	formats, err := video.Formats()
	if err != nil {
		t.Fatal(err)
	}

	if len(formats.Progressive) == 0 {
		t.Error("len(formats.Progressive) == 0")
	} else {
		if formats.Progressive.Best().URL == "" {
			t.Error("ProgressiveFormat.URL == \"\"")
		}
	}
	if formats.Hls.Url() == "" {
		t.Error("formats.Hls.Url() == \"\"")
	}
	if formats.Dash.Url() == "" {
		t.Error("formats.Dash.Url() == \"\"")
	}
}

func TestSearch(t *testing.T) {
	client := NewSearchClient()

	client.Filter = VideoFilter
	result, err := client.Search("Rick Astley", 1)
	if err != nil {
		t.Fatal(err)
	}
	videos := result.Data.Videos()
	if len(videos) == 0 {
		t.Fatal("len(videos) == 0")
	}
	if videos[0].Link == "" {
		t.Error("videos[0].Link == \"\"")
	}

	client.Filter = ChannelFilter
	result, err = client.Search("My channel", 1)
	if err != nil {
		t.Fatal(err)
	}
	channels := result.Data.Channels()
	if len(channels) == 0 {
		t.Fatal("len(channels) == 0")
	}
	if channels[0].Link == "" {
		t.Error("channels[0].Link == \"\"")
	}

	client.Filter = GroupFilter
	result, err = client.Search("Group", 1)
	if err != nil {
		t.Fatal(err)
	}
	groups := result.Data.Groups()
	if len(groups) == 0 {
		t.Fatal("len(groups) == 0")
	}
	if groups[0].Link == "" {
		t.Error("groups[0].Link == \"\"")
	}

	client.Filter = PeopleFilter
	result, err = client.Search("Mister", 1)
	if err != nil {
		t.Fatal(err)
	}
	people := result.Data.People()
	if len(people) == 0 {
		t.Fatal("len(people) == 0")
	}
	if people[0].Link == "" {
		t.Error("people[0].Link == \"\"")
	}
}
