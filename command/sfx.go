package command

import (
	"os/exec"
)

func Sfx(sfx string) {
	wavpath := "/home/bluecmd/go/src/github.com/hackverket/swedish-embassy-broadcasting/" + sfx + ".wav"

	paplayArgs := append([]string{
		"-s", "127.0.0.1",
		"--channel-map=aux0,aux1",
		wavpath,
	})

	exec.Command("paplay", paplayArgs...)
}
