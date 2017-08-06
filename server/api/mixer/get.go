package mixer

import (
  "time"
	"github.com/gin-gonic/gin"
  "github.com/hackverket/swedish-embassy-broadcasting/motuavb"
)

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
	c.JSON(200, level)
}
