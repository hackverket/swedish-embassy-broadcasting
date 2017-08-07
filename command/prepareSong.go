package command

import (
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/hackverket/swedish-embassy-broadcasting/mpd"
	uuid "github.com/satori/go.uuid"
)

func PrepareSong(filename string) {
	newPath := path.Join(os.Getenv("DUMP_PATH"), uuid.NewV4().String()) + ".opus"
	transcoding := exec.Command(
		"ffmpeg",
		"-f",
		"-i",
		filename,
		"-acodec",
		"flac",
		"-ar",
		"44100",
		"-ac",
		"2",
		newPath)
	o, err := transcoding.Output()

	log.Println(o, err)

	normalizing := exec.Command(
		"ffmpeg-normalize",
		"-f",
		"--level",
		"-18",
		"--format",
		"flac",
		"--no-prefix",
		newPath)

	f, berr := normalizing.Output()
	log.Println(f, berr)

	mpd.M.Add(newPath)
}
