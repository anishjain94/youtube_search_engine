package youtube

import (
	"context"
	"youtube_search_engine/common"
	youtubeIntegration "youtube_search_engine/integrations/youtube_integration"
)

func getVideosFromYoutube(ctx *context.Context) *[]YoutubeData {

	videoData := youtubeIntegration.SearchQuery(ctx, "surfing")
	print(videoData)

	videoModel := ToModel(videoData)
	createYoutubeData(ctx, videoModel)

	return videoModel
}

func getVideosFromDb(ctx *context.Context) *[]YoutubeData {

	query := common.GetQueryValueFromCtx(ctx)

	searchQuery := common.ParseToTsQuery(query.Get("query"))

	return getYoutubeData(ctx, searchQuery)
}
