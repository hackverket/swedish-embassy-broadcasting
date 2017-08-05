package motuavb

import (
	"fmt"
	"strconv"
	"time"

	tween "github.com/draoncc/go-tween"
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
	fmt.Println(c.GetFloat32Value(channelString))

	tween.NewEngine(time.Second, curves.Linear)
}
