package playlist

type Writer interface {
	WriteMedia(Media) error
	Flush() error
}
