package song

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hackverket/swedish-embassy-broadcasting/command"
)

func uploadSong(c *gin.Context) {
	//song := c.Request.FormValue("song")
	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		fmt.Println(err)
		return
	}
	filename := header.Filename
	fmt.Println(header.Filename)
	tempFilename := "/tmp/" + filename
	out, err := os.Create(tempFilename)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
		return
	}

	go command.PrepareSong(tempFilename)

	c.Redirect(http.StatusSeeOther, "/")
}
