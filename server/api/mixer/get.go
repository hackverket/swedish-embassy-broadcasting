package mixer

import (
  "time"
  "strconv"
  "log"
	"github.com/gin-gonic/gin"
  "github.com/hackverket/swedish-embassy-broadcasting/motuavb"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
  level []float64
)

func init() {
  go func() {
    etag := ""
    for {
      sc := motuavb.Connect("10.44.22.107")
      level, etag = sc.GetMeters(etag)
      time.Sleep(10 * time.Millisecond)
    }
  }()
}

func getMixer(c *gin.Context) {
  w := c.Writer
  r := c.Request
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

  for {
    // Channel 22 is one of the output channels
    conn.WriteMessage(websocket.TextMessage, []byte(strconv.FormatFloat(level[22], 'g', 1, 64)))
    time.Sleep(10 * time.Millisecond)
  }
}
