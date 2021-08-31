// Package vimego: Search, download Vimeo videos and retrieve metadata.
package vimego

import (
	"fmt"
	"net/http"
)

const UserAgent = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:84.0) Gecko/20100101 Firefox/84.0"

// NewVideo creates a new Video from URL.
func NewVideo(url string) (*Video, error) {
	videoId := validateUrl(url)
	if videoId == 0 {
		return nil, ErrInvalidUrl
	}

	return &Video{
		Url:        url,
		VideoId:    videoId,
		HTTPClient: &http.Client{},
		Header:     map[string][]string{"User-Agent": {UserAgent}},
	}, nil
}

// NewVideo creates a new Video from video ID.
func NewVideoFromId(videoId int) *Video {
	return &Video{
		Url:        fmt.Sprintf("https://vimeo.com/%v", videoId),
		VideoId:    videoId,
		HTTPClient: &http.Client{},
		Header:     map[string][]string{"User-Agent": {UserAgent}},
	}
}

// NewSearchClient creates a new SearchClient with default parameters.
func NewSearchClient() *SearchClient {
	return &SearchClient{
		PerPage:    18,
		Filter:     VideoFilter,
		Order:      RelevanceOrder,
		Direction:  DescDirection,
		Category:   AnyCategory,
		HTTPClient: &http.Client{},
		Header:     map[string][]string{"User-Agent": {UserAgent}},
	}
}
