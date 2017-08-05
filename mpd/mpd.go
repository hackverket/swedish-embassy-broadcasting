package mpd

import (
	"log"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/fhs/gompd/mpd"
)

var (
	M MpdClient
)

func init() {
	M.Host = "[::1]:6600"
	M.Init()
}

type MpdClient struct {
	Host string
}

func (c MpdClient) Init() {
	if os.Getenv("DISABLE_PURGE") == "" {
		go c.playlistPurge()
	}
}

func (c MpdClient) Add(f string) {
	conn, err := mpd.Dial("tcp", c.Host)
	if err != nil {
		log.Fatalln(err)
	}

	b := path.Base(f)
	os.Symlink(f, path.Join(os.Getenv("MPD_MUSIC_HOME"), b))

	log.Printf("Adding %s\n", b)
	conn.Update(b)
	// Yeah... I know. TODO and all that
	time.Sleep(1 * time.Second)
	conn.Add(b)
}

func (c MpdClient) playlistPurge() {
	conn, err := mpd.Dial("tcp", c.Host)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		time.Sleep(2 * time.Second)
		o, err := conn.CurrentSong()
		if err != nil {
			continue
		}
		id, err := strconv.Atoi(o["Id"])
		if err != nil {
			continue
		}
		if id == 1 {
			continue
		}
		log.Println("Purging old playlist entries")
		conn.Delete(id, id-1)
	}
}
