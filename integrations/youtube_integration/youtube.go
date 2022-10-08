package youtubeIntegration

import (
	"context"
	"net/http"
	"os"
	"sync"

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

func SearchQuery(ctx *context.Context, ch chan<- string, wg *sync.WaitGroup, query string) {

	wg.Add(1)

	youtubeClient := YoutubeClient.R()
	queryParams := "/youtube/v3/search?part=snippet&maxResults=10&q=" + query + "&key=" + YOUTUBE_KEY

	resp, err := youtubeClient.
		Get(queryParams)

	if err != nil || resp.StatusCode() != http.StatusOK {
		print(err.Error())
	}
	ch <- string(resp.Body())

}
