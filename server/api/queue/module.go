package queue

import (
	"github.com/gin-gonic/gin"
)

func New(parent *gin.RouterGroup) {
	parent.GET("/", getQueue)
}
