package command

import (
	"os/exec"

	"github.com/hackverket/swedish-embassy-broadcasting/motuavb"
	"github.com/hackverket/swedish-embassy-broadcasting/polly"
)

func TextToSpeech(text string) {

	ttsFile := polly.GetTTS(text)
	sc := motuavb.Connect("10.44.22.107")
	sc.FadeChannelVolume(8, 0.0)

	paplayArgs := append([]string{
		"-s", "127.0.0.1",
		"--channel-map=", "'aux0,aux1'",
		ttsFile,
	})
	exec.Command("paplay", paplayArgs...)
	sc.FadeChannelVolume(8, 0.7)
}
