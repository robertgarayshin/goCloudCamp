package main

import (
	"fmt"
	"goCloudCamp/pkg/track"
)

func main() {
	var inp int
	var pas string

	fmt.Scan(&inp)
	t := track.Track{Name: "aa", Duration: inp}
	go track.Track.Play(t)
	fmt.Scan(&pas)
	if pas == "p" {
		track.Track.Pause(t)
	}

}
