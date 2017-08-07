package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hackverket/swedish-embassy-broadcasting/server/api/effect"
	"github.com/hackverket/swedish-embassy-broadcasting/server/api/mixer"
	"github.com/hackverket/swedish-embassy-broadcasting/server/api/queue"
	"github.com/hackverket/swedish-embassy-broadcasting/server/api/song"
	"github.com/hackverket/swedish-embassy-broadcasting/server/api/tts"
)

func New(parent *gin.RouterGroup) {
	songGroup := parent.Group("/song")
	song.New(songGroup)
	queueGroup := parent.Group("/queue")
	queue.New(queueGroup)
	ttsGroup := parent.Group("/tts")
	tts.New(ttsGroup)
	mixerGroup := parent.Group("/mixer")
	mixer.New(mixerGroup)
	effectGroup := parent.Group("/effect")
	effect.New(effectGroup)
}
