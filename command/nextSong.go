package command

import (
	"github.com/hackverket/swedish-embassy-broadcasting/motuavb"
	"github.com/hackverket/swedish-embassy-broadcasting/mpd"
)

func NextSong() {

	sc := motuavb.Connect("10.44.22.107")
	sc.FadeChannelVolume(8, 0.05)

	mpd.M.Next()

	sc.FadeChannelVolume(8, 0.8)

}
