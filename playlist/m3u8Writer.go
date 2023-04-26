package playlist

import (
	"bufio"
	"fmt"
	"io"
)

type m3u8Writer struct {
	w *bufio.Writer
}

func (m3u8w *m3u8Writer) WriteMedia(media Media) (err error) {
	content := fmt.Sprintf("\r\n#EXTINF:000,%s\r\n%s\r\n", media.Title, media.Url)
	_, err = m3u8w.w.Write([]byte(content))
	return
}

func (m3u8w *m3u8Writer) Flush() error {
	return m3u8w.w.Flush()
}

func NewM3U8Writer(w io.Writer) (Writer, error) {
	bw := bufio.NewWriterSize(w, 15)
	_, err := bw.Write([]byte("#EXTM3U\r\n"))
	return &m3u8Writer{w: bw}, err
}
