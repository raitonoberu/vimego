package vimego

type VideoFormats struct {
	Progressive ProgressiveFormats `json:"progressive"`
	Dash        *DashFormat        `json:"dash"`
	Hls         *HlsFormat         `json:"hls"`
}

type ProgressiveFormats []*ProgressiveFormat

func (p ProgressiveFormats) Len() int {
	return len(p)
}

func (p ProgressiveFormats) Less(a, b int) bool {
	return p[a].Width < p[b].Width
}

func (p ProgressiveFormats) Swap(a, b int) {
	p[a], p[b] = p[b], p[a]
}

// Best returns the ProgressiveFormat with the hightest resolution.
func (p ProgressiveFormats) Best() *ProgressiveFormat {
	if len(p) != 0 {
		return p[len(p)-1]
	}
	return nil
}

// Worst returns the ProgressiveFormat with the lowest resolution.
func (p ProgressiveFormats) Worst() *ProgressiveFormat {
	if len(p) != 0 {
		return p[0]
	}
	return nil
}

type ProgressiveFormat struct {
	Profile int    `json:"profile"`
	Width   int    `json:"width"`
	Mime    string `json:"mime"`
	Fps     int    `json:"fps"`
	URL     string `json:"url"`
	Cdn     string `json:"cdn"`
	Quality string `json:"quality"`
	Origin  string `json:"origin"`
	Height  int    `json:"height"`
}

type DashFormat struct {
	SeparateAv bool   `json:"separate_av"`
	DefaultCdn string `json:"default_cdn"`
	Cdns       struct {
		AkfireInterconnectQuic struct {
			URL    string `json:"url"`
			Origin string `json:"origin"`
			AvcURL string `json:"avc_url"`
		} `json:"akfire_interconnect_quic"`
		FastlySkyfire struct {
			URL    string `json:"url"`
			Origin string `json:"origin"`
			AvcURL string `json:"avc_url"`
		} `json:"fastly_skyfire"`
	} `json:"cdns"`
}

// Url returns the URL for the video stream.
func (s *DashFormat) Url() string {
	switch s.DefaultCdn {
	case "akfire_interconnect_quic":
		return s.Cdns.AkfireInterconnectQuic.URL
	case "fastly_skyfire":
		return s.Cdns.FastlySkyfire.URL
	default:
		// fallback
		if s.Cdns.AkfireInterconnectQuic.URL != "" {
			return s.Cdns.AkfireInterconnectQuic.URL
		}
		if s.Cdns.FastlySkyfire.URL != "" {
			return s.Cdns.FastlySkyfire.URL
		}
	}
	return ""
}

type HlsFormat struct {
	SeparateAv bool   `json:"separate_av"`
	DefaultCdn string `json:"default_cdn"`
	Cdns       struct {
		AkfireInterconnectQuic struct {
			URL    string `json:"url"`
			Origin string `json:"origin"`
			AvcURL string `json:"avc_url"`
		} `json:"akfire_interconnect_quic"`
		FastlySkyfire struct {
			URL    string `json:"url"`
			Origin string `json:"origin"`
			AvcURL string `json:"avc_url"`
		} `json:"fastly_skyfire"`
	} `json:"cdns"`
}

// Url returns the URL for the .m3u8 playlist.
func (s *HlsFormat) Url() string {
	switch s.DefaultCdn {
	case "akfire_interconnect_quic":
		return s.Cdns.AkfireInterconnectQuic.URL
	case "fastly_skyfire":
		return s.Cdns.FastlySkyfire.URL
	default:
		// fallback
		if s.Cdns.AkfireInterconnectQuic.URL != "" {
			return s.Cdns.AkfireInterconnectQuic.URL
		}
		if s.Cdns.FastlySkyfire.URL != "" {
			return s.Cdns.FastlySkyfire.URL
		}
	}
	return ""
}
