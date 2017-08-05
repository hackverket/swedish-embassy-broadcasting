package song

import (
	"github.com/gin-gonic/gin"
	"github.com/hackverket/swedish-embassy-broadcasting/fetch"
	"github.com/hackverket/swedish-embassy-broadcasting/mpd"
)

type songRequest struct {
	URL string `form:"url" json:"url" binding:"required"`
}

func postSong(c *gin.Context) {
	var json songRequest
	if c.BindJSON(&json) == nil {

		go command.queueSong(json.URL)
		c.JSON(
			200,
			gin.H{"message": "Song is requested and added to the playlist."},
		)

	}
}

func processSong(url string) {
	au, _ := fetch.DownloadAudio(url)
	m := mpd.MpdClient{}
	m.Host = "[::1]:6600"
	m.Init()
	m.Add(au.Path)
}
