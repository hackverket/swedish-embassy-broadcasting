package command

import (
	"os"
	"os/exec"
	"path"

	"github.com/hackverket/swedish-embassy-broadcasting/mpd"
	uuid "github.com/satori/go.uuid"
)

func PrepareSong(filename string) {
	newPath := path.Join(os.Getenv("DUMP_PATH"), uuid.NewV4().String()) + ".opus"
	exec.Command(
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

	exec.Command(
		"ffmpeg-normalize",
		"-f",
		"--level",
		"-18",
		"--format",
		"flac",
		"--no-prefix",
		newPath)

	mpd.M.Add(newPath)
}
