package vimego

import (
	"fmt"
	"net/http"
)

const UserAgent = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:84.0) Gecko/20100101 Firefox/84.0"

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

func NewVideoFromId(videoId int) *Video {
	return &Video{
		Url:        fmt.Sprintf("https://vimeo.com/%v", videoId),
		VideoId:    videoId,
		HTTPClient: &http.Client{},
		Header:     map[string][]string{"User-Agent": {UserAgent}},
	}
}
