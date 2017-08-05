package queue

import (
	"github.com/gin-gonic/gin"
	"github.com/hackverket/swedish-embassy-broadcasting/mpd"
)

func getQueue(c *gin.Context) {
	c.JSON(200, mpd.M.GetQueue())
}
