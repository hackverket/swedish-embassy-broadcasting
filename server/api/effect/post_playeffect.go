package effect

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hackverket/swedish-embassy-broadcasting/command"
)

type effectRequest struct {
	effectName string `form:"name" json:"name" binding:"required"`
}

func playEffect(c *gin.Context) {
	var json effectRequest
	if c.BindJSON(&json) == nil {

		fmt.Println(json.effectName)
		go command.Sfx(json.effectName)
		c.JSON(
			200,
			gin.H{"message": "Effect will play."},
		)

	}
}
