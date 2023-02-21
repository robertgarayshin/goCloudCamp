package track

import (
	"fmt"
	"time"
)

func (t Track) Play() {
	for i := 0; i < t.Duration; {
		fmt.Println("Playing")
		time.Sleep(1000 * time.Millisecond)
		i += 1
	}
}

func (t Track) Pause() {
	fmt.Println("Paused")
}
