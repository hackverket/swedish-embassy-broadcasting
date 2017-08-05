package motuavb

import (
	"fmt"
	"time"

	tween "github.com/draoncc/go-tween"
)

type Fade struct {
	StartVolume float64
	EndVolume   float64
	Updates     chan float64
	Done        chan int
}

func NewFade(start float64, end float64) *Fade {
	return &Fade{StartVolume: start, EndVolume: end}
}

func (f *Fade) Start(framerate, frames int, frameTime, runningTime time.Duration) {
	fmt.Println(framerate, frames, frameTime, runningTime)
}

func (f *Fade) Update(frame tween.Frame) {
	f.Updates <- 0.0
	fmt.Println(frame.Transitioned)
}

func (f *Fade) End() {
	close(f.Done)
	fmt.Println("HELLO")
}
