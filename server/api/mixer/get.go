package mixer

import (
	"github.com/gin-gonic/gin"
  "github.com/hackverket/swedish-embassy-broadcasting/motuavb"
)

func getMixer(c *gin.Context) {
  sc := motuavb.Connect("10.44.22.107")
	c.JSON(200, sc.GetMeters("mix/level/0"))
}
