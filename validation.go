package vimego

import (
	"regexp"
	"strconv"
)

var validationPatterns = []regexp.Regexp{
	*regexp.MustCompile(`^https://player.vimeo.com/video/(?P<id>\d+)$`),
	*regexp.MustCompile(`^https://vimeo.com/(?P<id>\d+)$`),
	*regexp.MustCompile(`^https://vimeo.com/groups/.+?/videos/(?P<id>\d+)$`),
	*regexp.MustCompile(`^https://vimeo.com/manage/videos/(?P<id>\d+)$`),
}

func validateUrl(url string) int {
	for _, pattern := range validationPatterns {
		match := pattern.FindStringSubmatch(url)
		if match != nil {
			id, err := strconv.ParseInt(match[1], 10, 32)
			if err == nil {
				return int(id)
			}
			panic(err)
		}
	}
	return 0
}
