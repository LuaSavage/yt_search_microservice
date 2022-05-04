package searchresult

import (
	"context"
	"encoding/json"
	"fmt"

	cache "github.com/LuaSavage/yt_search_microservice/pkg/client/cache"
	redis "github.com/go-redis/redis/v8"
)

type Storage interface {
	GetSearchResultByQuary(ctx context.Context, query string) (*SearchResult, error)
	CreateSearchResult(ctx context.Context, searchResult *SearchResult) error
}

type storage struct {
	client cache.Client
}

func NewStorage(client cache.Client) Storage {
	return &storage{client: client}
}

func (s *storage) GetSearchResultByQuary(ctx context.Context, query string) (*SearchResult, error) {

	var searchResult SearchResult

	// get search result hash by type:id from redis
	if err := s.client.HGetAll(ctx, "search_result:"+query).Scan(&searchResult); err != nil {
		return nil, err
	}

	return nil, fmt.Errorf("search results by query '%s' alrady exists", query)
}

// put it in a cache
func (s *storage) CreateSearchResult(ctx context.Context, searchResult *SearchResult) error {

	if _, err := s.GetSearchResultByQuary(ctx, searchResult.Query); err == nil {
		return err
	}

	// trying to write it into redis
	if _, err := s.client.Pipelined(ctx, func(rdb redis.Pipeliner) error {

		// dto contains slice of video.id in place  of video models
		searchResultDTO := NewStoreSearchResultDTO(*searchResult)
		rdb.HSet(ctx, "search_result:"+searchResultDTO.Query, "query", searchResultDTO.Query)

		data, _ := json.Marshal(searchResultDTO.Videos)
		rdb.HSet(ctx, "search_result:"+searchResultDTO.Query, "videos", data)

		return nil
	}); err != nil {
		return err
	}

	return nil
}
