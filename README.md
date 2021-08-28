# vimego

### Search, download Vimeo videos and retrieve metadata.

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

	// You can even get a direct URL (progressive / hls)
	// url := video.Files.Progressive().Best()

	jsonstr, _ := json.Marshal(video)
	fmt.Println(string(jsonstr))
}
```
<details>
 <summary>Example Result</summary>

```json
{
   "name":"The Rick Astley Remixer",
   "description":"Together with the talented students of Dalarna's University, Erik made this interactive, Arduino based music remixer.\n\nThe concept: In realtime you get to remix Rick Astley's legendary \"Never Gonna Give You Up\". You do it simply by pushing a bunch of buttons on a physical, custom built controller. \n   You can change all the elements of the song and for example make a nice mashup of black metal, rap, gospel, gameboy sounds and Spanish vocals. You can also change the intensity and mood of the mix. It's just about as weird and fun as it sounds.\n\nOn a more technical note: The Arduino controller is connected to Flash via Tinkerproxy. The flash application that plays all the music uses Dinahmoe's AS3 framework for the adaptive music control as well as the loading of almost 600mb of mp3 files.\n\n--------\n\nA group of eight students from Dalarna's University (Christoffer Johansson, Mattias Prada, Christian Rosenberg, Emil Hemming, Joakim Lilliemarck Jansson, Joakim OldÃ©en, Gustav Look, Johan Edbom) helped develop the concept, plan the project and pick all the musical styles to be used. They also put their heart and soul into producing all the amazing music used in the project! \n\nErik Brattlöf, www.twitter.com/erikbrattlof, created the original concept, and programmed and designed the Arduino controller. Fillippe Åhlund did all the soldering and put all electronics for the controller together.\n\nDon't hesitate to get in touch if you want to test it out or display it somewhere!\n\nwww.dinahmoe.com\nwww.twitter.com/DinahmoeSTHLM",
   "type":"video",
   "link":"https://vimeo.com/dinahmoe/the-rick-astley-project",
   "duration":182,
   "width":1280,
   "language":"",
   "height":720,
   "created_time":"2011-06-21T22:30:02Z",
   "modified_time":"2018-08-10T14:40:25Z",
   "release_time":"2011-06-21T22:30:02Z",
   "content_rating":[
      "unrated"
   ],
   "pictures":{
      "active":true,
      "sizes":[
         {
            "width":100,
            "height":75,
            "link":"https://i.vimeocdn.com/video/167407170_100x75?r=pad"
         },
         {
            "width":200,
            "height":150,
            "link":"https://i.vimeocdn.com/video/167407170_200x150?r=pad"
         },
         {
            "width":295,
            "height":166,
            "link":"https://i.vimeocdn.com/video/167407170_295x166?r=pad"
         },
         {
            "width":640,
            "height":360,
            "link":"https://i.vimeocdn.com/video/167407170_640x360?r=pad"
         },
         {
            "width":960,
            "height":540,
            "link":"https://i.vimeocdn.com/video/167407170_960x540?r=pad"
         },
         {
            "width":1280,
            "height":720,
            "link":"https://i.vimeocdn.com/video/167407170_1280x720?r=pad"
         },
         {
            "width":1920,
            "height":1080,
            "link":"https://i.vimeocdn.com/video/167407170_1920x1080?r=pad"
         }
      ],
      "default_picture":false
   },
   "tags":[],
   "uploader":{
      "pictures":{
         "active":true,
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
         ],
         "default_picture":false
      }
   },
   "metadata":{
      "connections":{
         "comments":{
            "total":2
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
      "created_time":"2011-04-25T13:44:38Z",
      "pictures":{
         "active":true,
         "type":"custom",
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
         ],
         "default_picture":false
      },
      "metadata":{
         "connections":{
            "albums":{
               "total":6
            },
            "channels":{
               "total":0
            },
            "followers":{
               "total":118
            },
            "following":{
               "total":5
            },
            "likes":{
               "total":8
            },
            "videos":{
               "total":77
            }
         }
      }
   },
   "files":[
      {
         "quality":"hd",
         "type":"video/mp4",
         "width":1280,
         "height":720,
         "expires":"2021-08-28T23:48:39Z",
         "link":"https://player.vimeo.com/play/55870780?s=25429948_1630162119_ece99eb79a9522ce89e9d925b7328e89\u0026sid=c253761ea60bb15a68abb57b5c0c2047de6cdf521630151319\u0026oauth2_token_id=",
         "created_time":"2011-06-21T22:41:34Z",
         "fps":25,
         "video_file_id":55870780,
         "size":59571510,
         "md5":"fcf7076a9cb19247f35d2cf2e145dafc",
         "public_name":"HD 720p",
         "size_short":"56.81MB"
      },
      {
         "quality":"sd",
         "type":"video/mp4",
         "width":640,
         "height":360,
         "expires":"2021-08-28T23:48:39Z",
         "link":"https://player.vimeo.com/play/55870342?s=25429948_1630162119_26e4a50054aa164be8d54585e160f3f4\u0026sid=c253761ea60bb15a68abb57b5c0c2047de6cdf521630151319\u0026oauth2_token_id=",
         "created_time":"2011-06-21T22:38:18Z",
         "fps":25,
         "video_file_id":55870342,
         "size":19552156,
         "md5":"e3e1202e1649d7c4d1a32d553d5e430e",
         "public_name":"SD",
         "size_short":"18.65MB"
      },
      {
         "quality":"mobile",
         "type":"video/mp4",
         "width":480,
         "height":272,
         "expires":"2021-08-28T23:48:39Z",
         "link":"https://player.vimeo.com/play/55869927?s=25429948_1630162119_f90bf8686319777ea61c783654f148eb\u0026sid=c253761ea60bb15a68abb57b5c0c2047de6cdf521630151319\u0026oauth2_token_id=",
         "created_time":"2011-06-21T22:34:56Z",
         "fps":25,
         "video_file_id":55869927,
         "size":8378950,
         "md5":"1911ef3bdfb1b70960fbf80970572eda",
         "public_name":"Mobile SD",
         "size_short":"7.99MB"
      },
      {
         "quality":"hls",
         "type":"video/mp4",
         "expires":"2021-08-28T13:48:39Z",
         "link":"https://player.vimeo.com/play/55870780,55870342,55869927/hls?s=25429948_1630158519_6dfa81bdd5b7e1d01cd36e41c8a58964\u0026context=Vimeo%5CController%5CApi%5CResources%5CSearchController.\u0026oauth2_token_id=",
         "created_time":"2011-06-21T22:41:34Z",
         "fps":25,
         "video_file_id":55870780,
         "size":59571510,
         "md5":"fcf7076a9cb19247f35d2cf2e145dafc",
         "public_name":"HD 720p",
         "size_short":"56.81MB",
         "link_secure":"https://player.vimeo.com/play/55870780,55870342,55869927/hls?s=25429948_1630158519_6dfa81bdd5b7e1d01cd36e41c8a58964\u0026context=Vimeo%5CController%5CApi%5CResources%5CSearchController.\u0026oauth2_token_id="
      }
   ],
   "status":"available",
   "is_playable":true,
   "has_audio":true
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

## License

**MIT License**, see [LICENSE](./LICENSE) file for additional information.
