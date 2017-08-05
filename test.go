package main

import (
	"sync"

	"github.com/hackverket/swedish-embassy-broadcasting/server"
	"github.com/hackverket/swedish-embassy-broadcasting/shitirc"
)

func main() {

	//sc := motuavb.Connect("10.44.22.107")
	//sc.FadeChannelVolume(8, 0.4)
	//shitirc.Connect("#sha2017", "irc.quakenet.org:6667", "CYKA2000")

	var wg sync.WaitGroup
	wg.Add(1)

	sirc := shitirc.Dial("#sha2017", "irc.quakenet.org:6667", "HACK4JESUS")
	go sirc.Connect()
	go server.Start()

	wg.Wait()
}
