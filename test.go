package main

import (
	"sync"

	"github.com/hackverket/swedish-embassy-broadcasting/server"
	"github.com/hackverket/swedish-embassy-broadcasting/shitirc"
)

func main() {

	//shitirc.Connect("#sha2017", "irc.quakenet.org:6667", "CYKA2000")

	var wg sync.WaitGroup
	wg.Add(1)

	sirc := shitirc.Dial("#sha2017", "irc.quakenet.org:6667", "HACK4JESUS")
	go sirc.Connect()
	go server.Start()

	wg.Wait()
}
