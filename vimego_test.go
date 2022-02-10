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
	if len(videos[0].Files) == 0 {
		t.Fatal("len(videos[0].Files) == 0")
	}
	if videos[0].Files.Best().Link == "" {
		t.Error("videos[0].Files.Best().Link == \"\"")
	}
}
