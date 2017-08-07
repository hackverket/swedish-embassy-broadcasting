package effect

import (
	"github.com/gin-gonic/gin"
)

func New(parent *gin.RouterGroup) {
	parent.POST("/play", playEffect)
}
