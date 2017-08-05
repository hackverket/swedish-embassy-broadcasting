package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hackverket/swedish-embassy-broadcasting/server/api/song"
)

func New(parent *gin.RouterGroup) {
	songGroup := parent.Group("/song")
	song.New(songGroup)
}
