# vimego

### Download Vimeo videos and retrieve metadata.

#### Largely based on [yashrathi](https://github.com/yashrathi-git)'s [vimeo_downloader](https://github.com/yashrathi-git/vimeo_downloader).

## Installing

```bash
go get github.com/raitonoberu/vimego
```

Please note that the API is not yet final and may change in the future.

## Usage

### Get a direct URL for the best available .mp4 stream (video+audio)

```go
package main

import (
	"fmt"
	"github.com/raitonoberu/vimego"
)

func main() {
	video, _ := vimego.NewVideo("https://vimeo.com/206152466")
	formats, err := video.Formats()
	if err != nil {
		panic(err)
	}

	fmt.Println(formats.Progressive.Best().URL)
	// https://vod-progressive.akamaized.net/...
}
```

### Get metadata

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/raitonoberu/vimego"
)

func main() {
	video, _ := vimego.NewVideo("https://vimeo.com/206152466")
	metadata, err := video.Metadata()
	if err != nil {
		panic(err)
	}

	jsonstr, _ := json.Marshal(metadata)
	fmt.Println(string(jsonstr))
}
```
<details>
 <summary>Example Result</summary>

```json
{
  "id": 206152466,
  "title": "Crystal Castles - Kept",
  "description": "",
  "url": "https://vimeo.com/206152466",
  "upload_date": "2017-02-28 18:07:25",
  "thumbnail_small": "http://i.vimeocdn.com/video/621091880_100x75",
  "thumbnail_medium": "http://i.vimeocdn.com/video/621091880_200x150",
  "thumbnail_large": "http://i.vimeocdn.com/video/621091880_640",
  "user_id": 19229427,
  "user_name": "Vladislav Donets",
  "user_url": "https://vimeo.com/donec",
  "user_portrait_small": "http://i.vimeocdn.com/portrait/8438592_30x30",
  "user_portrait_medium": "http://i.vimeocdn.com/portrait/8438592_75x75",
  "user_portrait_large": "http://i.vimeocdn.com/portrait/8438592_100x100",
  "user_portrait_huge": "http://i.vimeocdn.com/portrait/8438592_300x300",
  "stats_number_of_likes": 211,
  "stats_number_of_plays": 65095,
  "stats_number_of_comments": 17,
  "duration": 243,
  "width": 1280,
  "height": 720,
  "tags": "Crystal Castles",
  "embed_privacy": "anywhere"
}
```
</details>

## Advanced usage

### About formats

Vimeo stores its streams in 3 different formats:
- **[Progressive](https://en.wikipedia.org/wiki/Progressive_download)**
    - A direct URL to the .mp4 stream (video+audio).
    - Max quality - 1080p.
    - Most probably, this is what you're looking for.
- **[Hls](https://en.wikipedia.org/wiki/HTTP_Live_Streaming)**
    - An URL to the master.m3u8 playlist.
    - Max quality - 2160p.
    - Best for passing it to video players (VLC, mpv, etc.).
- **[Dash](https://en.wikipedia.org/wiki/Dynamic_Adaptive_Streaming_over_HTTP)**
    - An URL to the JSON containing data about segments.
    - Max quality - 2160p.
    - Suitable if you need a video-only or audio-only stream.

### Get video-only or audio-only stream

There is a `Video.GetDashStreams` method that parses the DASH format and provides information about the available streams.

```go
package main

import (
	"io"
	"os"
	"github.com/raitonoberu/vimego"
)

func main() {
	video, _ := vimego.NewVideo("https://vimeo.com/206152466")
	formats, _ := video.Formats()
	streams, _ := video.GetDashStreams(formats.Dash.Url())

	stream, _, _ := streams.Audio.Best().Reader(nil) // io.ReadCloser
	file, _ := os.Create("output.m4a")
	defer file.Close()
	io.Copy(file, stream)
}
```

### Get embed-only videos

If the video you want to download can only be played on a specific site, there is a way to get its streams. You need to set the value `Referer` in the headers. Note that `Video.Metadata()` does not work with such videos.

```go
package main

import (
	"fmt"
	"github.com/raitonoberu/vimego"
)

func main() {
	video, _ := vimego.NewVideo("https://player.vimeo.com/video/498617513")
	video.Header["Referer"] = []string{"https://atpstar.com/plans-162.html"}

	formats, _ := video.Formats()
	fmt.Println(formats.Progressive.Best().URL)
}

```

## Information

The code seems to be ready, but I have some thoughts on improving it and there are still formalities to be completed.

### TODO:
- Add docs & comments
- Try to simplify the API
- Add example file
- Add tests
- Make a CLI tool
- Implement searching

## License

**MIT License**, see [LICENSE](./LICENSE) file for additional information.
