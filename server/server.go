package server

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/hackverket/swedish-embassy-broadcasting/server/api"
)

func Start() {
	r := gin.Default()

	sAPI := r.Group("/api")
	api.New(sAPI)

	r.Use(static.Serve("/", static.LocalFile("/home/bluecmd/go/src/github.com/hackverket/swedish-embassy-broadcasting/server/public", true)))

	r.Run(":4020")
}
