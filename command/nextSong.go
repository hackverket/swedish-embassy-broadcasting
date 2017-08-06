package command

import (
	"github.com/hackverket/swedish-embassy-broadcasting/mpd"
)

func NextSong() {

	mpd.M.Next()
}
