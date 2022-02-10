# vimego

### Search, download Vimeo videos and retrieve metadata.

#### Largely based on [yashrathi](https://github.com/yashrathi-git)'s [vimeo_downloader](https://github.com/yashrathi-git/vimeo_downloader).

## Installing

```bash
go get github.com/raitonoberu/vimego
```

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

### Search for videos

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/raitonoberu/vimego"
)

func main() {
	client := vimego.NewSearchClient()
	result, err := client.Search("Rick Astley", 1)
	if err != nil {
		panic(err)
	}
	video := result.Data.Videos()[0]

	jsonstr, _ := json.Marshal(video)
	fmt.Println(string(jsonstr))
}
```
<details>
 <summary>Example Result</summary>

```json
{
   "name":"The Rick Astley Remixer",
   "link":"https://vimeo.com/dinahmoe/the-rick-astley-project",
   "duration":182,
   "created_time":"2011-06-21T22:30:02Z",
   "privacy":{
      "view":"anybody"
   },
   "pictures":{
      "sizes":[
         {
            "width":100,
            "height":75,
            "link":"https://i.vimeocdn.com/video/167407170-345a400d1c7c4919f9bf098da33dba5673eb0cba165da8559516acd3e64d7f07-d_100x75?r=pad"
         },
         {
            "width":200,
            "height":150,
            "link":"https://i.vimeocdn.com/video/167407170-345a400d1c7c4919f9bf098da33dba5673eb0cba165da8559516acd3e64d7f07-d_200x150?r=pad"
         },
         {
            "width":295,
            "height":166,
            "link":"https://i.vimeocdn.com/video/167407170-345a400d1c7c4919f9bf098da33dba5673eb0cba165da8559516acd3e64d7f07-d_295x166?r=pad"
         },
         {
            "width":640,
            "height":360,
            "link":"https://i.vimeocdn.com/video/167407170-345a400d1c7c4919f9bf098da33dba5673eb0cba165da8559516acd3e64d7f07-d_640x360?r=pad"
         },
         {
            "width":960,
            "height":540,
            "link":"https://i.vimeocdn.com/video/167407170-345a400d1c7c4919f9bf098da33dba5673eb0cba165da8559516acd3e64d7f07-d_960x540?r=pad"
         },
         {
            "width":1280,
            "height":720,
            "link":"https://i.vimeocdn.com/video/167407170-345a400d1c7c4919f9bf098da33dba5673eb0cba165da8559516acd3e64d7f07-d_1280x720?r=pad"
         },
         {
            "width":1920,
            "height":1080,
            "link":"https://i.vimeocdn.com/video/167407170-345a400d1c7c4919f9bf098da33dba5673eb0cba165da8559516acd3e64d7f07-d_1920x1080?r=pad"
         }
      ]
   },
   "metadata":{
      "connections":{
         "comments":{
            "total":3
         },
         "likes":{
            "total":32
         }
      }
   },
   "user":{
      "name":"DinahmoeSTHLM",
      "link":"https://vimeo.com/dinahmoe",
      "location":"Stockholm, Sweden",
      "pictures":{
         "sizes":[
            {
               "width":30,
               "height":30,
               "link":"https://i.vimeocdn.com/portrait/17506926_30x30"
            },
            {
               "width":72,
               "height":72,
               "link":"https://i.vimeocdn.com/portrait/17506926_72x72"
            },
            {
               "width":75,
               "height":75,
               "link":"https://i.vimeocdn.com/portrait/17506926_75x75"
            },
            {
               "width":100,
               "height":100,
               "link":"https://i.vimeocdn.com/portrait/17506926_100x100"
            },
            {
               "width":144,
               "height":144,
               "link":"https://i.vimeocdn.com/portrait/17506926_144x144"
            },
            {
               "width":216,
               "height":216,
               "link":"https://i.vimeocdn.com/portrait/17506926_216x216"
            },
            {
               "width":288,
               "height":288,
               "link":"https://i.vimeocdn.com/portrait/17506926_288x288"
            },
            {
               "width":300,
               "height":300,
               "link":"https://i.vimeocdn.com/portrait/17506926_300x300"
            },
            {
               "width":360,
               "height":360,
               "link":"https://i.vimeocdn.com/portrait/17506926_360x360"
            }
         ]
      }
   }
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

The code seems to be ready, but I have some thoughts on improving it.

### TODO:
- Handle video IDs other than int
- Captcha processing
- Make a CLI tool

## License

**MIT License**, see [LICENSE](./LICENSE) file for additional information.
