package command

import (
	"fmt"
	"os/exec"
)

func Sfx(sfx string) {

	fmt.Println(sfx)
	wavpath := "/home/bluecmd/go/src/github.com/hackverket/swedish-embassy-broadcasting/sounds/" + sfx + ".wav"

	paplayArgs := append([]string{
		"-s", "127.0.0.1",
		"--channel-map=aux2,aux3",
		wavpath,
	})

	cmd := exec.Command("paplay", paplayArgs...)
	o, err := cmd.Output()
	fmt.Println(o, err)
}
