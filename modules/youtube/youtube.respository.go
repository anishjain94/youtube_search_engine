package youtube

import (
	"context"
	"youtube_search_engine/common"
	"youtube_search_engine/infra/database"

	"gorm.io/gorm"
)

func createYoutubeData(ctx *context.Context, data *[]YoutubeData) {

	db := database.GetDb(ctx)

	result := db.Save(data)
	if result.Error != nil {
		print(result.Error)
	}

}

func getYoutubeData(ctx *context.Context, searchQuery string) *[]YoutubeData {

	db := database.GetDb(ctx)
	youtubeData := &[]YoutubeData{}

	var result *gorm.DB

	if searchQuery != "" {
		result = db.Scopes(common.PaginatedScope(ctx)).Find(youtubeData, "to_tsquery(?) @@ to_tsvector(title || description)", searchQuery)
	} else {
		result = db.Scopes(common.PaginatedScope(ctx)).Find(youtubeData)
	}

	if result.Error != nil {
		print(result.Error)
	}
	return youtubeData
}
