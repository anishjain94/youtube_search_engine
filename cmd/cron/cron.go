package main

import (
	"context"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Seconds().Do(func(ctx context.Context) {
		// print(ctx)
		print("running cron")
	}, context.Background())

	go s.StartBlocking()

	<-context.Background().Done()
}
