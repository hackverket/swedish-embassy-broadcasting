package shitirc

import (
	"fmt"
	"regexp"
	"time"

	humanize "github.com/dustin/go-humanize"
	"github.com/hackverket/swedish-embassy-broadcasting/command"
	"github.com/hackverket/swedish-embassy-broadcasting/mpd"
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
	irccon := irc.IRC(c.nickname, "SWEDISHEMBASSY")
	irccon.VerboseCallbackHandler = false
	irccon.Debug = false
	irccon.AddCallback("001", func(e *irc.Event) { irccon.Join(c.channel) })
	irccon.AddCallback("366", func(e *irc.Event) {})
	irccon.AddCallback("PRIVMSG", func(event *irc.Event) {
		go func(event *irc.Event) {
			r := regexp.MustCompile(`\!r (.*)`)
			if r.MatchString(event.Message()) {

				go c.ircQueueSong(r.FindAllStringSubmatch(event.Message(), -1)[0][1], irccon)
			}

			swedish := regexp.MustCompile(`\!s (.*)`)
			if swedish.MatchString(event.Message()) {

				go command.TextToSpeech(swedish.FindAllStringSubmatch(event.Message(), -1)[0][1], "Astrid")
			}

			english := regexp.MustCompile(`\!ve (.*)`)
			if english.MatchString(event.Message()) {

				go command.TextToSpeech(english.FindAllStringSubmatch(event.Message(), -1)[0][1], "Joey")
			}

			german := regexp.MustCompile(`\!vg (.*)`)
			if german.MatchString(event.Message()) {

				go command.TextToSpeech(german.FindAllStringSubmatch(event.Message(), -1)[0][1], "Hans")
			}

			japanese := regexp.MustCompile(`\!vj (.*)`)
			if japanese.MatchString(event.Message()) {

				go command.TextToSpeech(japanese.FindAllStringSubmatch(event.Message(), -1)[0][1], "Mizuki")
			}

			finnish := regexp.MustCompile(`\!vf (.*)`)
			if finnish.MatchString(event.Message()) {

				go command.TextToSpeech(finnish.FindAllStringSubmatch(event.Message(), -1)[0][1], "Carmen")
			}

			a := regexp.MustCompile(`\!a (.*)`)
			if a.MatchString(event.Message()) {

				go command.Sfx(a.FindAllStringSubmatch(event.Message(), -1)[0][1])
			}

			n := regexp.MustCompile(`\!npppp`)
			if n.MatchString(event.Message()) {
				go command.NextSong()
			}

			//go polly.GetTTS(event.Message(), "", "message"+strconv.Itoa(k)+".mp3")
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
	go c.printQueue(irccon)
	irccon.Loop()
}

func (c *Client) printQueue(irccon *irc.Connection) {
	title := ""

	for {
		time.Sleep(1 * time.Second)
		new_title := mpd.M.GetQueue()[0].Title
		if title != new_title {
			irccon.Privmsg(c.channel, "Now playing:"+new_title)
		}
		title = new_title
	}
}

func (c *Client) ircQueueSong(url string, irccon *irc.Connection) {
	command.QueueSong(url)

	// What's a race condition? I don't know  ¯\_(ツ)_/¯
	queue := mpd.M.GetQueue()
	i := len(queue) - 1
	playing := time.Now().Add(time.Duration(queue[i].Duration) * time.Second)
	irccon.Privmsg(c.channel, "Queued "+queue[i].Title+" (in "+humanize.Time(playing)+")")
}
