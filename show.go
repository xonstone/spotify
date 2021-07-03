package spotify

import (
	"net/url"
	"strconv"
)

type SavedShow struct {
	// The date and time the show was saved, represented as an ISO
	// 8601 UTC timestamp with a zero offset (YYYY-MM-DDTHH:MM:SSZ).
	// You can use the TimestampLayout constant to convert this to
	// a time.Time value.
	AddedAt    string `json:"added_at"`
	SimpleShow `json:"show"`
}

// FullShow contains full data about a show.
type FullShow struct {
	SimpleShow

	// A list of the show’s episodes.
	Episodes SimpleEpisodePage `json:"episodes"`
}

// SimpleShow contains basic data about a show.
type SimpleShow struct {
	// A list of the countries in which the show can be played,
	// identified by their ISO 3166-1 alpha-2 code.
	AvailableMarkets []string `json:"available_markets"`

	// The copyright statements of the show.
	Copyrights []Copyright `json:"copyrights"`

	// A description of the show.
	Description string `json:"description"`

	// Whether or not the show has explicit content
	// (true = yes it does; false = no it does not OR unknown).
	Explicit bool `json:"explicit"`

	// Known external URLs for this show.
	ExternalURLs map[string]string `json:"external_urls"`

	// A link to the Web API endpoint providing full details
	// of the show.
	Endpoint string `json:"href"`

	// The SpotifyID for the show.
	ID ID `json:"id"`

	// The cover art for the show in various sizes,
	// widest first.
	Images []Image `json:"images"`

	// True if all of the show’s episodes are hosted outside
	// of Spotify’s CDN. This field might be null in some cases.
	IsExternallyHosted *bool `json:"is_externally_hosted"`

	// A list of the languages used in the show, identified by
	// their ISO 639 code.
	Languages []string `json:"languages"`

	// The media type of the show.
	MediaType string `json:"media_type"`

	// The name of the show.
	Name string `json:"name"`

	// The publisher of the show.
	Publisher string `json:"publisher"`

	// The object type: “show”.
	Type string `json:"type"`

	// The Spotify URI for the show.
	URI URI `json:"uri"`
}

// GetShow retrieves information about a specific show.
// API reference: https://developer.spotify.com/documentation/web-api/reference/#endpoint-get-a-show
func (c *Client) GetShow(id string) (*FullShow, error) {
	return c.GetShowOpt(nil, id)
}

// GetShowOpt is like GetShow while supporting an optional market parameter.
// API reference: https://developer.spotify.com/documentation/web-api/reference/#endpoint-get-a-show
func (c *Client) GetShowOpt(opt *Options, id string) (*FullShow, error) {
	spotifyURL := c.baseURL + "shows/" + id
	if opt != nil {
		v := url.Values{}
		if opt.Country != nil {
			v.Set("market", *opt.Country)
		}
		if params := v.Encode(); params != "" {
			spotifyURL += "?" + params
		}
	}

	var result FullShow

	err := c.get(spotifyURL, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// GetShowEpisodes retrieves paginated episode information about a specific show.
// API reference: https://developer.spotify.com/documentation/web-api/reference/#endpoint-get-a-shows-episodes
func (c *Client) GetShowEpisodes(id string) (*SimpleEpisodePage, error) {
	return c.GetShowEpisodesOpt(nil, id)
}

// GetShowEpisodesOpt is like GetShowEpisodes while supporting optional market, limit, offset parameters.
// API reference: https://developer.spotify.com/documentation/web-api/reference/#endpoint-get-a-shows-episodes
func (c *Client) GetShowEpisodesOpt(opt *Options, id string) (*SimpleEpisodePage, error) {
	spotifyURL := c.baseURL + "shows/" + id + "/episodes"
	if opt != nil {
		v := url.Values{}
		if opt.Country != nil {
			v.Set("market", *opt.Country)
		}
		if opt.Limit != nil {
			v.Set("limit", strconv.Itoa(*opt.Limit))
		}
		if opt.Offset != nil {
			v.Set("offset", strconv.Itoa(*opt.Offset))
		}
		if params := v.Encode(); params != "" {
			spotifyURL += "?" + params
		}
	}

	var result SimpleEpisodePage

	err := c.get(spotifyURL, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
