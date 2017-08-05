package tts

import (
	"github.com/gin-gonic/gin"
	"github.com/hackverket/swedish-embassy-broadcasting/command"
)

type songRequest struct {
	Text string `form:"text" json:"text" binding:"required"`
}

func postSong(c *gin.Context) {
	var json songRequest
	if c.BindJSON(&json) == nil {

		go command.TextToSpeech(json.Text)
		c.JSON(
			200,
			gin.H{"message": "TTS will play."},
		)

	}
}
