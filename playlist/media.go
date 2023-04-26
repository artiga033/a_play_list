package playlist

import "time"

type Media struct {
	Url   string
	Size  int64
	Title string
	Date  *time.Time
	Thumb string
}
