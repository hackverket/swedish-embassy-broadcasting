package motuavb

import (
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
	return &Fade{
		StartVolume: start,
		EndVolume:   end,
		Updates:     make(chan float64),
		Done:        make(chan int),
	}
}

func (f *Fade) Start(framerate, frames int, frameTime, runningTime time.Duration) {

}

func (f *Fade) Update(frame tween.Frame) {
	d := f.EndVolume - f.StartVolume
	f.Updates <- d*frame.Transitioned + f.StartVolume

}

func (f *Fade) End() {
	close(f.Done)
}
