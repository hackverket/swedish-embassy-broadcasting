package song

import (
	"github.com/gin-gonic/gin"
	"github.com/hackverket/swedish-embassy-broadcasting/command"
)

type songRequest struct {
	URL string `form:"url" json:"url" binding:"required"`
}

func postSong(c *gin.Context) {
	var json songRequest
	if c.BindJSON(&json) == nil {

		go command.QueueSong(json.URL)
		c.JSON(
			200,
			gin.H{"message": "Song is requested and added to the playlist."},
		)

	}
}
