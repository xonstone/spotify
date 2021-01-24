package spotify

import "encoding/json"

type SumSimpleEpisodeSimpleTrack struct {
	SimpleEpisode *SimpleEpisode
	*SimpleTrack
}

func (p *SumSimpleEpisodeSimpleTrack) UnmarshalJSON(data []byte) error {
	i := &typedItem{}

	if err := json.Unmarshal(data, i); err != nil {
		return err
	}

	switch i.Type {
	case "episode":
		p.SimpleEpisode = &SimpleEpisode{}
		if err := json.Unmarshal(data, p.SimpleEpisode); err != nil {
			return err
		}
	case "track":
		fallthrough
	default:
		p.SimpleTrack = &SimpleTrack{}
		if err := json.Unmarshal(data, p.SimpleTrack); err != nil {
			return err
		}
	}

	return nil

}

type SumFullEpisodeFullTrack struct {
	FullEpisode *FullEpisode
	*FullTrack
}

func (p *SumFullEpisodeFullTrack) UnmarshalJSON(data []byte) error {
	i := &typedItem{}

	if err := json.Unmarshal(data, i); err != nil {
		return err
	}

	switch i.Type {
	case "episode":
		p.FullEpisode = &FullEpisode{}
		if err := json.Unmarshal(data, p.FullEpisode); err != nil {
			return err
		}
	case "track":
		fallthrough
	default:
		p.FullTrack = &FullTrack{}
		if err := json.Unmarshal(data, p.FullTrack); err != nil {
			return err
		}
	}

	return nil

}

type typedItem struct {
	Type string `json:"type"`
}
