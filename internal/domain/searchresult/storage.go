package searchresult

import (
	"context"
	"encoding/json"
	"fmt"

	cache "github.com/LuaSavage/yt_search_microservice/pkg/client/cache"
	redis "github.com/go-redis/redis/v8"
)

type Storage interface {
	GetSearchResultByQuary(ctx context.Context, query string) (*StoreSearchResultDTO, error)
	CreateSearchResult(ctx context.Context, searchResult *SearchResult) error
}

type storage struct {
	client cache.Client
}

func NewStorage(client cache.Client) Storage {
	return &storage{client: client}
}

func (s *storage) GetSearchResultByQuary(ctx context.Context, query string) (*StoreSearchResultDTO, error) {
	var searchResultDTO StoreSearchResultDTO

	// get search result hash by type:id from redis
	if err := s.client.HGetAll(ctx, "search_result:"+query).Scan(&searchResultDTO); err != nil {
		return nil, err
	}

	if searchResultDTO.Query != query {
		return nil, fmt.Errorf("search results by query '%s' does'nt exists", query)
	}

	return &searchResultDTO, nil
}

// put it in a cache
func (s *storage) CreateSearchResult(ctx context.Context, searchResult *SearchResult) error {

	if _, err := s.GetSearchResultByQuary(ctx, searchResult.Query); err == nil {
		return fmt.Errorf("search result bu queary '%s' already exists in cache", searchResult.Query)
	}

	// trying to write it into redis
	_, err := s.client.Pipelined(ctx, func(rdb redis.Pipeliner) error {

		// dto contains slice of video.id in place  of video models
		searchResultDTO := NewStoreSearchResultDTO(*searchResult)
		cmd := rdb.HSet(ctx, "search_result:"+searchResultDTO.Query, "query", searchResultDTO.Query)

		if cmd.Err() != nil {
			return cmd.Err()
		}

		data, _ := json.Marshal(searchResultDTO.Videos)
		cmd = rdb.HSet(ctx, "search_result:"+searchResultDTO.Query, "videos", data)

		if cmd.Err() != nil {
			return cmd.Err()
		}

		return nil
	})

	return err
}
