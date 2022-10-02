package youtube

import youtubeIntegration "youtube_search_engine/integrations/youtube_integration"

func ToModel(videoDataDto *youtubeIntegration.YoutubeVideoDto) *[]YoutubeData {

	youtubeModel := []YoutubeData{}

	for _, videoData := range videoDataDto.Items {

		youtube := YoutubeData{
			Title:        videoData.Sinppet.Title,
			Description:  videoData.Sinppet.Description,
			ChannelTitle: videoData.Sinppet.ChannelTitle,
			PublishedAt:  videoData.Sinppet.PublishedAt,
			Url:          "https://www.youtube.com/watch?v=" + videoData.Id.VideoId,
		}
		youtubeModel = append(youtubeModel, youtube)
	}

	return &youtubeModel
}
