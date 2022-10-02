package youtube

import (
	"context"
	youtubeIntegration "youtube_search_engine/integrations/youtube_integration"
)

func getVideosFromYoutube(ctx *context.Context) *[]YoutubeData {

	videoData := youtubeIntegration.SearchQuery(ctx, "surfing")
	print(videoData)

	videoModel := ToModel(videoData)
	createYoutubeData(ctx, videoModel)

	return videoModel
}
