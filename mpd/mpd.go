package mpd

import (
  "fmt"
  "github.com/fhs/gompd/mpd"
  "log"
  "os"
  "path"
  "strconv"
  "time"
)

type MpdClient struct {
  Host string
  conn *mpd.Client
}

func (c MpdClient) Init() { conn, err := mpd.Dial("tcp", c.Host)
  if err != nil {
    log.Fatalln(err)
  }
  c.conn = conn
  if os.Getenv("DISABLE_PURGE") == "" {
    go c.playlistPurge()
  }
}

func (c MpdClient) Add(f string) {
  b := path.Base(f)
  os.Symlink(f, path.Join(os.Getenv("MPD_MUSIC_HOME"), b))

  c.conn.Update(b)
  c.conn.Add(b)
}

func (c MpdClient) playlistPurge() {
  for {
    time.Sleep(2 * time.Second)
    o, err := c.conn.CurrentSong()
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
    fmt.Println("Purging old playlist entries")
    c.conn.Delete(id, id-1)
  }
}
