package vimego

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type SearchClient struct {
	PerPage   int
	Filter    SearchFilter
	Order     SortOrder
	Direction SortDirection
	Category  SearchCategory

	Header     map[string][]string
	HTTPClient *http.Client

	token   string
	tokenMu sync.Mutex
}

func (c *SearchClient) getToken() (string, error) {
	req, _ := http.NewRequest("GET", "https://vimeo.com/_rv/jwt", nil)
	req.Header = map[string][]string{"X-Requested-With": {"XMLHttpRequest"}}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Token string `json:"token"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", ErrDecodingFailed
	}

	return result.Token, nil
}

// Search returns the result from the requested page.
func (c *SearchClient) Search(query string, page int) (*SearchResult, error) {
	var token string
	c.tokenMu.Lock()
	if c.token == "" {
		newToken, err := c.getToken()
		if err != nil {
			c.tokenMu.Unlock()
			return nil, err
		}
		token = newToken
		c.token = newToken
	} else {
		token = c.token
	}
	c.tokenMu.Unlock()

	params := url.Values{}
	params.Add("fields", "search_web")
	params.Add("query", query)
	params.Add("filter_type", string(c.Filter))
	params.Add("sort", string(c.Order))
	params.Add("direction", string(c.Direction))
	if c.Category != "" {
		params.Add("filter_category", string(c.Category))
	}
	params.Add("page", fmt.Sprint(page))
	params.Add("per_page", fmt.Sprint(c.PerPage))
	req, _ := http.NewRequest("GET", "https://api.vimeo.com/search?"+params.Encode(), nil)
	req.Header = c.Header
	req.Header["Authorization"] = []string{"jwt " + token}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result := SearchResult{}

	if resp.StatusCode == 200 {
		err := json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			return nil, ErrDecodingFailed
		}
	} else {
		if resp.StatusCode == 401 {
			// trying to received a new token
			c.token = ""
			var token string
			c.tokenMu.Lock()
			if c.token == "" {
				newToken, err := c.getToken()
				if err != nil {
					c.tokenMu.Unlock()
					return nil, err
				}
				token = newToken
				c.token = newToken
			} else {
				// the token was received in parallel request
				// while the mutex was locked
				token = c.token
			}
			c.tokenMu.Unlock()

			req.Header["Authorization"] = []string{"jwt " + token}
			resp, err := c.HTTPClient.Do(req)
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				return nil, ErrUnexpectedStatusCode(resp.StatusCode)
			}

			err = json.NewDecoder(resp.Body).Decode(&result)
			if err != nil {
				return nil, ErrDecodingFailed
			}
		} else {
			return nil, ErrUnexpectedStatusCode(resp.StatusCode)
		}
	}

	return &result, nil
}

type SearchResult struct {
	Total   int        `json:"total"`
	Page    int        `json:"page"`
	PerPage int        `json:"per_page"`
	Data    SearchData `json:"data"`
}

type SearchData []struct {
	Type    string       `json:"type"`
	Video   *VideoItem   `json:"clip,omitempty"`
	People  *PeopleItem  `json:"people,omitempty"`
	Channel *ChannelItem `json:"channel,omitempty"`
	Group   *GroupItem   `json:"group,omitempty"`
}

func (d SearchData) Videos() []*VideoItem {
	result := []*VideoItem{}
	for _, item := range d {
		if item.Type == "clip" {
			result = append(result, item.Video)
		}
	}
	return result
}

func (d SearchData) People() []*PeopleItem {
	result := []*PeopleItem{}
	for _, item := range d {
		if item.Type == "people" {
			result = append(result, item.People)
		}
	}
	return result
}

func (d SearchData) Channels() []*ChannelItem {
	result := []*ChannelItem{}
	for _, item := range d {
		if item.Type == "channel" {
			result = append(result, item.Channel)
		}
	}
	return result
}

func (d SearchData) Groups() []*GroupItem {
	result := []*GroupItem{}
	for _, item := range d {
		if item.Type == "group" {
			result = append(result, item.Group)
		}
	}
	return result
}

type VideoItem struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Type          string    `json:"type"`
	Link          string    `json:"link"`
	Duration      int       `json:"duration"`
	Width         int       `json:"width"`
	Language      string    `json:"language"`
	Height        int       `json:"height"`
	CreatedTime   time.Time `json:"created_time"`
	ModifiedTime  time.Time `json:"modified_time"`
	ReleaseTime   time.Time `json:"release_time"`
	ContentRating []string  `json:"content_rating"`
	Pictures      struct {
		Active         bool           `json:"active"`
		Sizes          []*PictureSize `json:"sizes"`
		DefaultPicture bool           `json:"default_picture"`
	} `json:"pictures"`
	Tags []struct {
		Name      string `json:"name"`
		Tag       string `json:"tag"`
		Canonical string `json:"canonical"`
		Metadata  struct {
			Connections struct {
				Videos struct {
					Total int `json:"total"`
				} `json:"videos"`
			} `json:"connections"`
		} `json:"metadata"`
	} `json:"tags"`
	Uploader struct {
		Pictures struct {
			Active         bool           `json:"active"`
			Sizes          []*PictureSize `json:"sizes"`
			DefaultPicture bool           `json:"default_picture"`
		} `json:"pictures"`
	} `json:"uploader"`
	Metadata struct {
		Connections struct {
			Comments struct {
				Total int `json:"total"`
			} `json:"comments"`
			Likes struct {
				Total int `json:"total"`
			} `json:"likes"`
		} `json:"connections"`
	} `json:"metadata"`
	User struct {
		Name        string    `json:"name"`
		Link        string    `json:"link"`
		Location    string    `json:"location"`
		CreatedTime time.Time `json:"created_time"`
		Pictures    struct {
			Active         bool           `json:"active"`
			Type           string         `json:"type"`
			Sizes          []*PictureSize `json:"sizes"`
			DefaultPicture bool           `json:"default_picture"`
		} `json:"pictures"`
		Metadata struct {
			Connections struct {
				Albums struct {
					Total int `json:"total"`
				} `json:"albums"`
				Channels struct {
					Total int `json:"total"`
				} `json:"channels"`
				Followers struct {
					Total int `json:"total"`
				} `json:"followers"`
				Following struct {
					Total int `json:"total"`
				} `json:"following"`
				Likes struct {
					Total int `json:"total"`
				} `json:"likes"`
				Videos struct {
					Total int `json:"total"`
				} `json:"videos"`
			} `json:"connections"`
		} `json:"metadata"`
	} `json:"user"`
	Files      VideoFiles `json:"files"`
	Status     string     `json:"status"`
	IsPlayable bool       `json:"is_playable"`
	HasAudio   bool       `json:"has_audio"`
}

type PictureSize struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Link   string `json:"link"`
}

type VideoFile struct {
	Quality     string    `json:"quality"`
	Type        string    `json:"type"`
	Width       int       `json:"width,omitempty"`
	Height      int       `json:"height,omitempty"`
	Expires     time.Time `json:"expires"`
	Link        string    `json:"link"`
	CreatedTime time.Time `json:"created_time"`
	Fps         float64   `json:"fps"`
	VideoFileID int       `json:"video_file_id"`
	Size        int       `json:"size"`
	Md5         string    `json:"md5"`
	PublicName  string    `json:"public_name"`
	SizeShort   string    `json:"size_short"`
	LinkSecure  string    `json:"link_secure,omitempty"`
}

type VideoFiles []*VideoFile

// Best returns the VideoFile with the largest size.
func (f VideoFiles) Best() *VideoFile {
	if len(f) == 0 {
		return nil
	}
	var best = f[0]

	for _, file := range f {
		if file.Size > best.Size {
			best = file
		}
	}
	return best
}

// Worst returns the VideoFile with the smallest size.
func (f VideoFiles) Worst() *VideoFile {
	if len(f) == 0 {
		return nil
	}
	var worst = f[0]

	for _, file := range f {
		if file.Size < worst.Size {
			worst = file
		}
	}
	return worst
}

// Prorgessive returns VideoFiles with a direct URL.
func (f VideoFiles) Progressive() VideoFiles {
	var result = VideoFiles{}
	for _, file := range f {
		if file.Quality != "hls" {
			result = append(result, file)
		}
	}
	return result
}

// Hls returns VideoFiles with a URL to the .m3u8 playlist.
func (f VideoFiles) Hls() *VideoFile {
	for _, file := range f {
		if file.Quality == "hls" {
			return file
		}
	}
	return nil
}

type PeopleItem struct {
	Name     string `json:"name"`
	Link     string `json:"link"`
	Location string `json:"location"`
	Pictures struct {
		Sizes []*PictureSize `json:"sizes"`
	} `json:"pictures"`
	Metadata struct {
		Connections struct {
			Followers struct {
				Total int `json:"total"`
			} `json:"followers"`
			Videos struct {
				Total int `json:"total"`
			} `json:"videos"`
		} `json:"connections"`
	} `json:"badge"`
}

type ChannelItem struct {
	Name     string `json:"name"`
	Link     string `json:"link"`
	Pictures struct {
		Sizes []struct {
			Width              int    `json:"width"`
			Height             int    `json:"height"`
			Link               string `json:"link"`
			LinkWithPlayButton string `json:"link_with_play_button"`
		} `json:"sizes"`
	} `json:"pictures"`
	Metadata struct {
		Connections struct {
			Users struct {
				Total int `json:"total"`
			} `json:"users"`
			Videos struct {
				Total int `json:"total"`
			} `json:"videos"`
		} `json:"connections"`
	} `json:"metadata"`
}

type GroupItem struct {
	Name     string `json:"name"`
	Link     string `json:"link"`
	Pictures struct {
		Sizes []struct {
			Width              int    `json:"width"`
			Height             int    `json:"height"`
			Link               string `json:"link"`
			LinkWithPlayButton string `json:"link_with_play_button"`
		} `json:"sizes"`
	} `json:"pictures"`
	Metadata struct {
		Connections struct {
			Users struct {
				Total int `json:"total"`
			} `json:"users"`
			Videos struct {
				Total int `json:"total"`
			} `json:"videos"`
		} `json:"connections"`
	} `json:"metadata"`
}
