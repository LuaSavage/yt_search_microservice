package searchresult

import (
	video "github.com/LuaSavage/yt_search_microservice/internal/domain/video"
	ytsearch "github.com/LuaSavage/yt_search_microservice/pkg/client/ytsearch"
	ytvideo "github.com/LuaSavage/yt_search_microservice/pkg/client/ytvideo"
)

type NewServiceDTO struct {
	SearchApi     ytsearch.Client
	Storage       Storage
	VideoService  video.Service
	YtVideoClient ytvideo.Client
}

type StoreSearchResultDTO struct {
	Query  string   `json:"query" yaml:"query"`
	Videos []string `json:"videos" yaml:"videos"`
}

func NewStoreSearchResultDTO(searchResult *SearchResult) *StoreSearchResultDTO {
	var videoIds []string

	result := StoreSearchResultDTO{Query: searchResult.Query}

	for _, video := range searchResult.Videos {
		videoIds = append(videoIds, video.Id)
	}

	result.Videos = videoIds
	return &result
}
