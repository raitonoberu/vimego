package vimego

import "time"

type Metadata struct {
	ID                 int    `json:"id"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	URL                string `json:"url"`
	UploadDate         string `json:"upload_date"`
	ThumbnailSmall     string `json:"thumbnail_small"`
	ThumbnailMedium    string `json:"thumbnail_medium"`
	ThumbnailLarge     string `json:"thumbnail_large"`
	UserID             int    `json:"user_id"`
	UserName           string `json:"user_name"`
	UserURL            string `json:"user_url"`
	UserPortraitSmall  string `json:"user_portrait_small"`
	UserPortraitMedium string `json:"user_portrait_medium"`
	UserPortraitLarge  string `json:"user_portrait_large"`
	UserPortraitHuge   string `json:"user_portrait_huge"`
	Likes              int    `json:"stats_number_of_likes"`
	Plays              int    `json:"stats_number_of_plays"`
	Comments           int    `json:"stats_number_of_comments"`
	Duration           int    `json:"duration"`
	Width              int    `json:"width"`
	Height             int    `json:"height"`
	Tags               string `json:"tags"`
	EmbedPrivacy       string `json:"embed_privacy"`
}

func (m *Metadata) GetUploadDate() (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", m.UploadDate)
}
