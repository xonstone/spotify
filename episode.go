package spotify

import (
	"strconv"
	"strings"
	"time"
)

// SimpleEpisode contains basic data about an episode.
type SimpleEpisode struct {
	// A URL to a 30 second preview (MP3 format) of the episode.
	AudioPreviewURL string `json:"audio_preview_url"`

	// A description of the episode.
	Description string `json:"description"`

	// The episode length in milliseconds.
	Duration int `json:"duration_ms"`

	// Whether or not the episode has explicit content
	// (true = yes it does; false = no it does not OR unknown).
	Explicit bool `json:"explicit"`

	// 	External URLs for this episode.
	ExternalURLs map[string]string `json:"external_urls"`

	// A link to the Web API endpoint providing full details of the episode.
	Endpoint string `json:"href"`

	// The Spotify ID for the episode.
	ID ID `json:"id"`

	// The cover art for the episode in various sizes, widest first.
	Images []Image `json:"images"`

	// True if the episode is hosted outside of Spotify’s CDN.
	IsExternallyHosted bool `json:"is_externally_hosted"`

	// True if the episode is playable in the given market.
	// Otherwise false.
	IsPlayable bool `json:"is_playable"`

	// A list of the languages used in the episode, identified by their ISO 639 code.
	Languages []string `json:"languages"`

	// The name of the episode.
	Name string `json:"name"`

	// The date the episode was first released, for example
	// "1981-12-15". Depending on the precision, it might
	// be shown as "1981" or "1981-12".
	ReleaseDate string `json:"release_date"`

	// The precision with which release_date value is known:
	// "year", "month", or "day".
	ReleaseDatePrecision string `json:"release_date_precision"`

	// The user’s most recent position in the episode. Set if the
	// supplied access token is a user token and has the scope
	// user-read-playback-position.
	ResumePoint ResumePointObject `json:"resume_point"`

	// The show on which the episode belongs.
	Show SimpleShow `json:"show"`

	// The object type: "episode".
	Type string `json:"type"`

	// The Spotify URI for the episode.
	URI URI `json:"uri"`
}

type FullEpisode struct {
	SimpleEpisode
}

type ResumePointObject struct {
	// 	Whether or not the episode has been fully played by the user.
	FullyPlayed bool `json:"fully_played"`

	// The user’s most recent position in the episode in milliseconds.
	ResumePositionMS int `json:"resume_position_ms"`
}

// ReleaseDateTime converts the show's ReleaseDate to a time.TimeValue.
// All of the fields in the result may not be valid.  For example, if
// ReleaseDatePrecision is "month", then only the month and year
// (but not the day) of the result are valid.
func (e *SimpleEpisode) ReleaseDateTime() time.Time {
	if e.ReleaseDatePrecision == "day" {
		result, _ := time.Parse(DateLayout, e.ReleaseDate)
		return result
	}
	if e.ReleaseDatePrecision == "month" {
		ym := strings.Split(e.ReleaseDate, "-")
		year, _ := strconv.Atoi(ym[0])
		month, _ := strconv.Atoi(ym[1])
		return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	}
	year, _ := strconv.Atoi(e.ReleaseDate)
	return time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
}
