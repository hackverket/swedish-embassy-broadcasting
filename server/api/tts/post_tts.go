package tts

import (
	"github.com/gin-gonic/gin"
	"github.com/hackverket/swedish-embassy-broadcasting/command"
)

type ttsRequest struct {
	Text string `form:"text" json:"text" binding:"required"`
}

func postTTS(c *gin.Context) {
	var json ttsRequest
	if c.BindJSON(&json) == nil {

		go command.TextToSpeech(json.Text)
		c.JSON(
			200,
			gin.H{"message": "TTS will play."},
		)

	}
}
