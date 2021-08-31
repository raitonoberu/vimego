package vimego

import (
	"bytes"
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
)

type DashStreams struct {
	ClipID  string           `json:"clip_id"`
	BaseURL string           `json:"base_url"`
	Video   DashVideoStreams `json:"video"`
	Audio   DashAudioStreams `json:"audio"`
}

type DashSegment struct {
	Start float64 `json:"start"`
	End   float64 `json:"end"`
	URL   string  `json:"url"`
	Size  int     `json:"size"`
}

type DashVideoStreams []*DashVideoStream

func (d DashVideoStreams) Len() int {
	return len(d)
}

func (d DashVideoStreams) Less(a, b int) bool {
	return d[a].Bitrate < d[b].Bitrate
}

func (d DashVideoStreams) Swap(a, b int) {
	d[a], d[b] = d[b], d[a]
}

// Best returns the DashVideoStream with the highest bitrate.
func (d DashVideoStreams) Best() *DashVideoStream {
	if len(d) != 0 {
		return d[len(d)-1]
	}
	return nil
}

// Worst returns the DashVideoStream with the lowest bitrate.
func (d DashVideoStreams) Worst() *DashVideoStream {
	if len(d) != 0 {
		return d[0]
	}
	return nil
}

type DashAudioStreams []*DashAudioStream

func (d DashAudioStreams) Len() int {
	return len(d)
}

func (d DashAudioStreams) Less(a, b int) bool {
	return d[a].Bitrate < d[b].Bitrate
}

func (d DashAudioStreams) Swap(a, b int) {
	d[a], d[b] = d[b], d[a]
}

// Best returns the DashAudioStream with the highest bitrate.
func (d DashAudioStreams) Best() *DashAudioStream {
	if len(d) != 0 {
		return d[len(d)-1]
	}
	return nil
}

// Worst returns the DashAudioStream with the lowest bitrate.
func (d DashAudioStreams) Worst() *DashAudioStream {
	if len(d) != 0 {
		return d[0]
	}
	return nil
}

type DashStream struct {
	ID                 string         `json:"id"`
	URL                string         `json:"url"`
	BaseURL            string         `json:"base_url"`
	Format             string         `json:"format"`
	MimeType           string         `json:"mime_type"`
	Codecs             string         `json:"codecs"`
	Bitrate            int            `json:"bitrate"`
	AvgBitrate         int            `json:"avg_bitrate"`
	Duration           float64        `json:"duration"`
	MaxSegmentDuration int            `json:"max_segment_duration"`
	InitSegment        string         `json:"init_segment"`
	Segments           []*DashSegment `json:"segments"`
}

// Readers returns an io.ReadCloser for reading streaming data.
func (s *DashStream) Reader(httpClient *http.Client) (io.ReadCloser, int64, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	req, _ := http.NewRequest("GET", "", nil)

	r, w := io.Pipe()
	var length int64

	initSegment, err := base64.StdEncoding.DecodeString(s.InitSegment)
	if err != nil {
		return nil, 0, err
	}

	length += int64(len(initSegment))
	for _, chunk := range s.Segments {
		length += int64(chunk.Size)
	}

	loadChunk := func(chunkUrl string) error {
		newUrl, err := url.Parse(chunkUrl)
		if err != nil {
			return err
		}
		req.URL = newUrl

		resp, err := httpClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 400 {
			return ErrUnexpectedStatusCode(resp.StatusCode)
		}

		_, err = io.Copy(w, resp.Body)
		return err
	}

	go func() {
		// load the init chunk
		_, err := io.Copy(w, bytes.NewReader(initSegment))
		if err != nil {
			_ = w.CloseWithError(err)
			return
		}

		// load all the chunks
		for _, chunk := range s.Segments {
			err := loadChunk(s.URL + chunk.URL)
			if err != nil {
				_ = w.CloseWithError(err)
				return
			}
		}

		w.Close()
	}()

	return r, length, nil
}

type DashVideoStream struct {
	Framerate float64 `json:"framerate"`
	Width     int     `json:"width"`
	Height    int     `json:"height"`
	DashStream
}

type DashAudioStream struct {
	Channels   int `json:"channels"`
	SampleRate int `json:"sample_rate"`
	DashStream
}
