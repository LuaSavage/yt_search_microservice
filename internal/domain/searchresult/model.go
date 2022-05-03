package searchresult

import (
	video "github.com/LuaSavage/yt_search_microservice/internal/domain/video"
)

type SearchResult struct {
	Query  string        `json:"query"`
	Videos []video.Video `json:"videos"`
}
