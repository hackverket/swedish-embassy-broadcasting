package polly

import (
	"io/ioutil"

	"github.com/leprosus/golang-tts"
)

func GetTTS(text string, voice string, file string) {

	polly := golang_tts.New("", "")
	polly.Format(golang_tts.MP3)
	polly.Voice(golang_tts.Astrid)

	bytes, err := polly.Speech(text)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("./"+file, bytes, 0644)
}
