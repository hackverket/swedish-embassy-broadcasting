package fetch

import (
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path"
)

type AudioOutput struct {
	Path      string
	Info      string
	CmdOutput string
}

func DownloadAudio(id string) (AudioOutput, error) {
	prefix := path.Join(os.Getenv("DUMP_PATH"), hex.EncodeToString([]byte(id)))
	cmd := exec.Command(
		"youtube-dl",
		"--extract-audio",
		"--audio-format=flac",
		"--write-info-json",
		"--max-filesize=30m",
		"--output="+prefix+".%(ext)s",
		id)
	o, err := cmd.Output()
	output := AudioOutput{
		Path:      prefix + ".flac",
		Info:      prefix + ".info.json",
		CmdOutput: string(o)}

	cmd2 := exec.Command(
		"ffmpeg-normalize",
		"-f",
		"--level",
		"-12",
		"--format",
		"flac",
		"--no-prefix",
		output.Path)
	l, berr := cmd2.Output()
	fmt.Println(l)

	return output, err
}
