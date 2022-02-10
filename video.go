package vimego

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strings"
)

type Video struct {
	Url     string
	VideoId int

	Header     map[string][]string
	HTTPClient *http.Client
}

// Metadata returns the video metadata.
func (v *Video) Metadata() (*Metadata, error) {
	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("http://vimeo.com/api/v2/video/%v.json",
			v.VideoId),
		nil,
	)
	req.Header = v.Header
	resp, err := v.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, ErrUnexpectedStatusCode(resp.StatusCode)
	}

	var result []*Metadata
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result[0], nil
}

// Formats returns the video formats.
// Progressive formats contain direct video+audio streams up to 1080p.
// Hls format contains an URL to .m3u8 playlist with all possible streams.
// Dash format contains a JSON URL that can be parsed using GetDashStreams.
func (v *Video) Formats() (*VideoFormats, error) {
	configUrl := fmt.Sprintf("https://player.vimeo.com/video/%v/config", v.VideoId)
	req, _ := http.NewRequest("GET", configUrl, nil)
	req.Header = v.Header
	resp, err := v.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var configData struct {
		Request struct {
			Files *VideoFormats `json:"files"`
		} `json:"request"`
	}

	if resp.StatusCode < 400 {
		err = json.NewDecoder(resp.Body).Decode(&configData)
		if err != nil {
			return nil, fmt.Errorf("couldn't decode config JSON: %w", err)
		}
	} else {
		if resp.StatusCode == 403 {
			// If the response is forbidden it tries another way to fetch link
			req, _ := http.NewRequest("GET", v.Url, nil)
			req.Header = v.Header
			resp, err := v.HTTPClient.Do(req)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()

			if resp.StatusCode < 400 {
				pattern := fmt.Sprintf(
					`"(%s.+?)"`,
					strings.ReplaceAll(configUrl, "/", `\\/`),
				)
				rexp, err := regexp.Compile(pattern)
				if err != nil {
					return nil, ErrParsingFailed
				}
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return nil, ErrParsingFailed
				}
				configUrls := rexp.FindAll(body, 1)
				if len(configUrls) == 0 {
					return nil, ErrParsingFailed
				}
				configUrl := strings.Trim(strings.ReplaceAll(string(configUrls[0]), `\/`, "/"), `"`)
				req, err := http.NewRequest("GET", configUrl, nil)
				if err != nil {
					return nil, ErrParsingFailed
				}
				req.Header = v.Header
				resp, err := v.HTTPClient.Do(req)
				if err != nil {
					return nil, err
				}
				defer resp.Body.Close()
				err = json.NewDecoder(resp.Body).Decode(&configData)
				if err != nil {
					return nil, fmt.Errorf("couldn't decode config JSON: %w", err)
				}
			} else {
				return nil, ErrParsingFailed
			}
		} else {
			return nil, ErrParsingFailed
		}
	}
	if configData.Request.Files == nil {
		return nil, ErrParsingFailed
	}
	sort.Sort(configData.Request.Files.Progressive)
	return configData.Request.Files, nil
}

// GetDashStreams returns DASH streams of the video.
func (v *Video) GetDashStreams(dashUrl string) (*DashStreams, error) {
	req, _ := http.NewRequest("GET", dashUrl, nil)
	req.Header = v.Header
	resp, err := v.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result DashStreams
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("couldn't decode dash JSON: %w", err)
	}

	formaturl, _ := url.Parse(dashUrl)
	baseurl, _ := url.Parse(result.BaseURL)
	baseurl = formaturl.ResolveReference(baseurl)

	for _, stream := range result.Video {
		refurl, _ := url.Parse(stream.BaseURL)
		stream.URL = baseurl.ResolveReference(refurl).String()
	}
	for _, stream := range result.Audio {
		refurl, _ := url.Parse(stream.BaseURL)
		stream.URL = baseurl.ResolveReference(refurl).String()
	}

	sort.Sort(result.Video)
	sort.Sort(result.Audio)

	return &result, nil
}
