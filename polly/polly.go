package polly

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/leprosus/golang-tts"
	uuid "github.com/satori/go.uuid"
)

func GetTTS(text string) string {

	polly := golang_tts.New(os.Getenv("PKEY"), os.Getenv("PSECRET"))
	polly.Format(golang_tts.MP3)
	polly.Voice(golang_tts.Astrid)

	bytes, err := polly.Speech(text)
	if err != nil {
		panic(err)
	}

	filename := path.Join(os.Getenv("DUMP_PATH"), uuid.NewV4().String()) + ".mp3"
	ioutil.WriteFile(filename, bytes, 0644)
	return filename
}
