package song

import "github.com/gin-gonic/gin"

func getSongs(c *gin.Context) {
	c.JSON(200, "here are song")
}
