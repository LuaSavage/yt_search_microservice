package searchresult

import (
	video "github.com/LuaSavage/yt_search_microservice/internal/domain/video"
)

type SearchResult struct {
	Query  string        `json:"query" redis:"query" yaml:"query"`
	Videos []video.Video `json:"videos" redis:"videos" yaml:"videos"`
}
