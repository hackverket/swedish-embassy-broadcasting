package song

import (
	"fmt"

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
			gin.H{"message": "Song is requested and added to the playlist."},
		)

	}
}

func processSong(url string) {
	test, test2 := fetch.DownloadAudio(url)
	fmt.Println(test, test2)
}
