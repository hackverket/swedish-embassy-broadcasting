package sauna

import (
	"time"

	"github.com/hackverket/swedish-embassy-broadcasting/command"
)

func Start() {
	ticker := time.NewTicker(time.Duration(250) * time.Second)
	go func() {
		for _ = range ticker.C {
			command.ReadSauna()
		}
	}()
}
