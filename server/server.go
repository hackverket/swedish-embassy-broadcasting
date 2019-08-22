package server

import (
	"os"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/hackverket/swedish-embassy-broadcasting/server/api"
)

func Start() {
	r := gin.Default()

	sAPI := r.Group("/api")
	api.New(sAPI)

	home,_ := os.UserHomeDir()
	r.Use(static.Serve("/", static.LocalFile(home + "/go/src/github.com/hackverket/swedish-embassy-broadcasting/server/public", true)))

	r.Run(":4020")
}
