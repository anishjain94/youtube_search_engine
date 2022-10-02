package youtubeIntegration

import (
	"context"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
)

var (
	YoutubeClient *resty.Client
	YOUTUBE_KEY   string
)

func InitializeYoutube() {

	YOUTUBE_KEY = os.Getenv("YOUTUBE_KEY")

	YoutubeClient = resty.New().
		SetHeader("Content-Type", "application/json").
		SetBaseURL("https://youtube.googleapis.com")
}

func SearchQuery(ctx *context.Context, query string) *YoutubeVideoDto {
	youtubeClient := YoutubeClient.R()
	queryParams := "/youtube/v3/search?part=snippet&maxResults=20&q=" + query + "&key=" + YOUTUBE_KEY

	resultDto := YoutubeVideoDto{}
	resp, err := youtubeClient.
		SetResult(&resultDto).
		Get(queryParams)

	if err != nil || resp.StatusCode() != http.StatusOK {
		print(err.Error())
	}

	return &resultDto
}
