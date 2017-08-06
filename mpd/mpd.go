package mpd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
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

type Queue struct {
	Image string
	Title string
	Duration int
	Files []string
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

func (c MpdClient) Next() {
	conn, err := mpd.Dial("tcp", c.Host)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	err = conn.Next()
	if err != nil {
		log.Fatalln(err)
	}
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

  status, err := conn.Status()
	if err != nil {
		log.Fatalln(err)
	}

  offset, _ := strconv.ParseFloat(status["elapsed"], 64)
  offset = offset * -1

	for _, element := range attrs {
		f := element["file"]

		i := Queue{}
    i.Duration = int(offset)
    o, _ := strconv.ParseFloat(element["duration"], 64)
    offset += o

    f = path.Join(os.Getenv("MPD_MUSIC_HOME"), f)
    i.Files = append(i.Files, f)
		link, err := os.Readlink(f)
		if err == nil {
      i.Files = append(i.Files, link)
			infopath := strings.TrimSuffix(link, path.Ext(link)) + ".info.json"
			b, err := ioutil.ReadFile(infopath)
			if err == nil {
        i.Files = append(i.Files, infopath)
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
	for {
		time.Sleep(10 * time.Second)

    go func() {
      conn, err := mpd.Dial("tcp", c.Host)
      if err != nil {
        log.Fatalln(err)
      }
      defer conn.Close()

      o, err := conn.CurrentSong()
      if err != nil {
        return
      }
      pos, err := strconv.Atoi(o["Pos"])
      if err != nil {
        return
      }
      if pos == 0 {
        return
      }
      log.Printf("Purging old playlist entries %v\n", pos)
      queue := M.GetQueue()
      for _, q := range queue[0:pos] {
        for _, f := range q.Files {
          log.Printf("Removing orphaned file %v\n", f)
          os.Remove(f)
        }
      }
      conn.Delete(0, pos)
    }()
	}
}
