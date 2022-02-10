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
		return "", fmt.Errorf("couldn't decode token JSON: %w", err)
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
			return nil, fmt.Errorf("couldn't decode search JSON: %w", err)
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
				return nil, fmt.Errorf("couldn't decode search JSON: %w", err)
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
	Name        string    `json:"name"`
	Link        string    `json:"link"`
	Duration    int       `json:"duration"`
	CreatedTime time.Time `json:"created_time"`
	Privacy     struct {
		View string `json:"view"`
	} `json:"privacy"`
	Pictures struct {
		Sizes []PictureSize `json:"sizes"`
	} `json:"pictures"`
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
		Name     string `json:"name"`
		Link     string `json:"link"`
		Location string `json:"location"`
		Pictures struct {
			Sizes []PictureSize `json:"sizes"`
		} `json:"pictures"`
	} `json:"user"`
}

type PeopleItem struct {
	Name     string `json:"name"`
	Link     string `json:"link"`
	Location string `json:"location"`
	Pictures struct {
		Sizes []PictureSize `json:"sizes"`
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
		Sizes []PictureSize `json:"sizes"`
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
		Sizes []PictureSize `json:"sizes"`
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

type PictureSize struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Link   string `json:"link"`
}
