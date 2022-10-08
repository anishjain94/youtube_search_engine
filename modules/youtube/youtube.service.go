package youtube

import (
	"context"
	"encoding/json"
	"sync"
	"youtube_search_engine/common"
	youtubeIntegration "youtube_search_engine/integrations/youtube_integration"
)

func GetVideosFromYoutube(ctx *context.Context) {

	ch := make(chan string)
	var wg sync.WaitGroup

	go youtubeIntegration.SearchQuery(ctx, ch, &wg, "surfing")

	resp := <-ch

	go func() {
		wg.Wait()
		close(ch)
	}()

	videoData := youtubeIntegration.YoutubeVideoDto{}
	json.Unmarshal([]byte(resp), &videoData)

	videoModel := ToModel(&videoData)
	go createYoutubeData(ctx, videoModel)

}

func getVideosFromDb(ctx *context.Context) *[]YoutubeData {

	query := common.GetQueryValueFromCtx(ctx)

	searchQuery := common.ParseToTsQuery(query.Get("query"))

	return getYoutubeData(ctx, searchQuery)
}
