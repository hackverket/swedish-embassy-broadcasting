package command

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/hackverket/swedish-embassy-broadcasting/motuavb"
	"github.com/hackverket/swedish-embassy-broadcasting/polly"
	uuid "github.com/satori/go.uuid"
)

func TextToSpeech(text string) {

	ttsFile := polly.GetTTS(text)
	sc := motuavb.Connect("10.44.22.107")
	sc.FadeChannelVolume(8, 0.05)
	wavpath := path.Join(os.Getenv("DUMP_PATH"), uuid.NewV4().String()) + ".wav"
	lol := exec.Command("ffmpeg", "-i", ttsFile, "-ar", "44100", "-ac", "2", wavpath)
	g, berr := lol.Output()
	fmt.Println(g, berr)

	paplayArgs := append([]string{
		"-s", "127.0.0.1",
		"--channel-map=aux0,aux1",
		wavpath,
	})

	fmt.Println(paplayArgs)

	cmd := exec.Command("paplay", paplayArgs...)
	o, err := cmd.Output()
	fmt.Println(o, err)
	sc.FadeChannelVolume(8, 0.8)
}
