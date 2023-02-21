package playlist

import (
	"goCloudCamp/pkg/track"
)

type Playlist struct {
	Head       *track.Track
	Tail       *track.Track
	NowPlaying track.Track
}

type ListActions interface {
	AddTrack()
	NextTrack()
	PrevTrack()
}
