package song

import (
	"github.com/gin-gonic/gin"
)

func New(parent *gin.RouterGroup) {
	parent.GET("/", getSongs)
	parent.POST("/request", postSong)
}
