package youtubeIntegration

import "time"

type PageInfoDto struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

// TODO: Convert time fields to datetime.
// TODO: Skipping thumbnails for now.
type SnippetDto struct {
	PublishedAt          time.Time `json:"publishedAt"`
	ChannelId            string    `json:"channelId"`
	Title                string    `json:"title"`
	Description          string    `json:"description"`
	ChannelTitle         string    `json:"channelTitle"`
	LiveBroadcastContent string    `json:"liveBroadcastContent"`
	PublishTime          string    `json:"publishTime"`
}

type IdDto struct {
	Kind    string `json:"kind"`
	VideoId string `json:"videoId"`
}

type ItemsDto struct {
	Kind    string     `json:"kind"`
	Etag    string     `json:"etag"`
	Sinppet SnippetDto `json:"snippet"`
	Id      IdDto      `json:"id"`
}

type YoutubeVideoDto struct {
	Kind          string      `json:"kind"`
	Etag          string      `json:"etag"`
	NextPageToken string      `json:"nextPageToken"`
	RegionCode    string      `json:"regionCode"`
	PageInfo      PageInfoDto `json:"pageInfo"`
	Items         []ItemsDto  `json:"items"`
}
