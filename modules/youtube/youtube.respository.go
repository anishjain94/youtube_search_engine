package youtube

import (
	"context"
	"youtube_search_engine/infra/database"
)

func createYoutubeData(ctx *context.Context, data *[]YoutubeData) {

	db := database.GetDb(ctx)

	result := db.Save(data)
	if result.Error != nil {
		print(result.Error)
	}

}
