package shitirc

import (
	"fmt"
	"strconv"

	"github.com/swedish-embassy-broadcasting/polly"
	irc "github.com/thoj/go-ircevent"
)

type Client struct {
	channel  string
	server   string
	nickname string
}

func Dial(channel string, server string, nickname string) (c *Client) {
	return &Client{
		channel:  channel,
		server:   server,
		nickname: nickname,
	}
}

func (c *Client) Connect() {
	var k int = 0
	irccon := irc.IRC(c.nickname, "SWEDISHEMBASSY")
	irccon.VerboseCallbackHandler = false
	irccon.Debug = false
	irccon.AddCallback("001", func(e *irc.Event) { irccon.Join(c.channel) })
	irccon.AddCallback("366", func(e *irc.Event) {})
	irccon.AddCallback("PRIVMSG", func(event *irc.Event) {
		go func(event *irc.Event) {
			k++
			go polly.GetTTS(event.Message(), "", "message"+strconv.Itoa(k)+".mp3")
			//event.Message() contains the message
			//event.Nick Contains the sender
			//event.Arguments[0] Contains the channel
		}(event)
	})
	err := irccon.Connect(c.server)
	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}
	irccon.Loop()
}
