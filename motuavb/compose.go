package motuavb

import (
	"fmt"
	"strconv"
	"time"

	tween "github.com/draoncc/go-tween"
	"github.com/draoncc/go-tween/easing"
)

type Client struct {
	ip string
}

func Connect(newIp string) (c *Client) {
	return &Client{ip: newIp}
}

func (c *Client) SetChannelVolume(channel int, volume float64) {

	channelString := `mix/chan/` + strconv.Itoa(channel) + `/matrix/fader`
	c.SendFloat32(channelString, volume)
}

func (c *Client) FadeChannelVolume(channel int, volume float64) {
	channelString := `mix/chan/` + strconv.Itoa(channel) + `/matrix/fader`
	oldValue := float64(c.GetFloat32Value(channelString))

	updater := NewFade(oldValue, volume)
	fader := tween.NewEngine(time.Second, easing.QuadOut, updater)
	fader.Start()

	running := true
	for running {
		select {
		case t := <-updater.Updates:
			fmt.Println(t)
		case <-updater.Done:
			running = false
		}
	}

}
