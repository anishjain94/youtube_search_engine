package main

import (
	"context"
	"time"
	"youtube_search_engine/infra/database"
	"youtube_search_engine/infra/environment"
	youtube_integration "youtube_search_engine/integrations/youtube_integration"
	"youtube_search_engine/modules/youtube"

	"github.com/go-co-op/gocron"
)

func main() {

	environment.InitializeEnvs()
	database.InitializeGorm()
	youtube_integration.InitializeYoutube()

	s := gocron.NewScheduler(time.UTC)

	s.Every(10).Seconds().Do(func(ctx context.Context) {
		youtube.GetVideosFromYoutube(&ctx)
	}, context.Background())

	go s.StartBlocking()

	<-context.Background().Done()
}
