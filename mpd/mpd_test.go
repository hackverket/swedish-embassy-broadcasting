package mpd

import (
  "testing"
  "time"
)

func TestMpd(t *testing.T) {
  m := MpdClient{}
  m.Host = "[::1]:6600"
  m.Init()

  time.Sleep(5 * time.Second)
}

func TestMpdAdd(t *testing.T) {
  m := MpdClient{}
  m.Host = "[::1]:6600"
  m.Init()
  m.Add("/home/bluecmd/testar.mp3")
}
