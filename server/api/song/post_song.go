package song

import (
	"github.com/gin-gonic/gin"
	"github.com/hackverket/swedish-embassy-broadcasting/fetch"
)

type songRequest struct {
	URL string `form:"url" json:"url" binding:"required"`
}

func postSong(c *gin.Context) {
	var json songRequest
	if c.BindJSON(&json) == nil {

		go processSong(json.URL)
		c.JSON(
			200,
			gin.H{"message": "Song is being requested"},
		)

	}
}

func processSong(url string) {
	fetch.DownloadAudio(url)
}
