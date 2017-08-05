package mpd

import (
	"log"
	"os"
	"path"
	"strconv"
	"time"
	"encoding/json"
	"strings"
	"io/ioutil"

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

type Queue struct {
	Image string
	Title string
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
	defer conn.Close()

	b := path.Base(f)
	os.Symlink(f, path.Join(os.Getenv("MPD_MUSIC_HOME"), b))

	log.Printf("Adding %s\n", b)
	conn.Update(b)
	// Yeah... I know. TODO and all that
	time.Sleep(1 * time.Second)
	conn.Add(b)
}

func (c MpdClient) GetQueue() []Queue {
	conn, err := mpd.Dial("tcp", c.Host)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	var s []Queue
	attrs, err := conn.PlaylistInfo(-1, -1)
	if err != nil {
		log.Fatalln(err)
	}

	for _, element := range attrs {
		f := element["file"]

		i := Queue{}
		link, err := os.Readlink(path.Join(os.Getenv("MPD_MUSIC_HOME"), f))
		if err == nil {
			infopath := strings.TrimSuffix(link, path.Ext(link)) + ".info.json"
			b, err := ioutil.ReadFile(infopath)
			if err == nil {
				var m interface{}
				err := json.Unmarshal(b, &m)
				info := m.(map[string]interface{})
				if err == nil {
					i.Image = info["thumbnail"].(string)
					i.Title = info["fulltitle"].(string)
				}
			} else {
				log.Println(err)
			}
		}
		s = append(s, i)
	}
	return s
}

func (c MpdClient) playlistPurge() {
	conn, err := mpd.Dial("tcp", c.Host)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	for {
		time.Sleep(10 * time.Second)
		o, err := conn.CurrentSong()
		if err != nil {
			continue
		}
		pos, err := strconv.Atoi(o["Pos"])
		if err != nil {
			continue
		}
		if pos == 0 {
			continue
		}
		log.Printf("Purging old playlist entries %v\n", pos)
		conn.Delete(0, pos)
	}
}