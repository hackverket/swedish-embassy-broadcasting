package effect

import (
	"github.com/gin-gonic/gin"
	"github.com/hackverket/swedish-embassy-broadcasting/command"
)

type effectRequest struct {
	EffectName string `form:"name" json:"name" binding:"required"`
}

func playEffect(c *gin.Context) {
	var json effectRequest
	if c.BindJSON(&json) == nil {

		go command.Sfx(json.EffectName)
		c.JSON(
			200,
			gin.H{"message": "Effect will play."},
		)

	}
}
