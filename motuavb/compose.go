package motuavb

import (
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
	fader := tween.NewEngine(time.Second*2, easing.QuadOut, updater)
	fader.Start()

	running := true
	for running {
		select {
		case t := <-updater.Updates:
			c.SetChannelVolume(channel, t)
		case <-updater.Done:
			running = false
		}
	}

}
