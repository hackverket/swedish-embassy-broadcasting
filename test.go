package main

import (
	"github.com/hackverket/swedish-embassy-broadcasting/motuavb"
)

func main() {

	sc := motuavb.Connect("10.44.22.107")
	sc.FadeChannelVolume(8, 0.5)
	//shitirc.Connect("#sha2017", "irc.quakenet.org:6667", "CYKA2000")

}
