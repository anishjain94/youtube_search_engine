package youtube

import (
	"time"

	"gorm.io/gorm"
)

type YoutubeData struct {
	gorm.Model
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	ChannelTitle string    `json:"channel_title"`
	Url          string    `json:"url"`
	PublishedAt  time.Time `json:"published_at"`
}
