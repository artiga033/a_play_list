package a_play_list

import (
	"io"
	"net/url"
	"strings"

	stdpath "path"

	"github.com/artiga033/a_play_list/alist"
	"github.com/artiga033/a_play_list/playlist"
)

type playlistGenerator struct {
	client alist.Client
	writer playlist.Writer
}

func NewGenerator(client alist.Client, writer playlist.Writer) *playlistGenerator {
	return &playlistGenerator{client: client, writer: writer}
}

func NewM3U8Generator(client alist.Client, w io.Writer) (*playlistGenerator, error) {
	writer, err := playlist.NewM3U8Writer(w)
	return &playlistGenerator{client: client, writer: writer}, err
}

func (gen *playlistGenerator) GenerateFor(path, password string) error {
	c := gen.client
	resp, err := c.FsList(alist.ListReq{
		PageReq:  alist.PageReq{Page: 1, PerPage: 0},
		Path:     path,
		Password: password,
	})
	if err != nil {
		return err
	}
	for _, item := range resp.Data.Content {
		ext := stdpath.Ext(item.Name)
		ext = strings.TrimPrefix(ext, ".")
		_, isVideo := VideoTypes[ext]
		if isVideo {
			link, _ := url.JoinPath(c.Endpoint, "d", path, item.Name)
			err = gen.writer.WriteMedia(
				playlist.Media{
					Title: item.Name,
					Url:   link,
					Thumb: item.Thumb,
					Size:  item.Size,
					Date:  &item.Modified,
				})
			if err != nil {
				return err
			}
		}
	}
	err = gen.writer.Flush()
	return err
}
