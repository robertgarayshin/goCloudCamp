package track

type Track struct {
	Name     string
	Author   string
	Duration int
	Prev     *Track
	Next     *Track
}

type Actions interface {
	Play(song Track)
	Pause(song Track)
}
