package playlist

import "github.com/robertgarayshin/goCloudCamp/song"

type Playlist struct {
	Song song.Song
	Prev *song.Song
	Next *song.Song
}
