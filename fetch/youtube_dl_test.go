package fetch

import "testing"

func TestDownload(t *testing.T) {

  t.Log(DownloadAudio("https://www.youtube.com/watch?v=dQw4w9WgXcQ"))
}
