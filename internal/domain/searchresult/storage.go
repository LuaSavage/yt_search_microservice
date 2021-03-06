package searchresult

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	cache "github.com/LuaSavage/yt_search_microservice/pkg/client/cache"
)

//mockery --name=Storage --filename=storage.go --output=../../mocks/searchresult/ --outpkg=searchmocks
type Storage interface {
	GetSearchResultByQuary(ctx context.Context, query string) (*StoreSearchResultDTO, error)
	CreateSearchResult(ctx context.Context, searchResult *SearchResult) error
}

type storage struct {
	client     cache.Client
	expiration int
}

func NewStorage(client cache.Client, expiration int) Storage {
	return &storage{
		client:     client,
		expiration: expiration,
	}
}

func (s *storage) expire(ctx context.Context, key string) error {
	if cmd := s.client.Expire(ctx, key, time.Second*time.Duration(s.expiration)); cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}

func (s *storage) GetSearchResultByQuary(ctx context.Context, query string) (*StoreSearchResultDTO, error) {
	var searchResultDTO *StoreSearchResultDTO = &StoreSearchResultDTO{}

	// get search result hash by type:id from redis
	stringMap := s.client.HGetAll(ctx, "search_result:"+query)

	if stringMap.Err() != nil {
		return nil, stringMap.Err()
	}

	searchResultDTO.Query = stringMap.Val()["query"]
	var ids []string = []string{}
	json.Unmarshal([]byte(stringMap.Val()["videos"]), &ids)
	searchResultDTO.Videos = ids

	if searchResultDTO.Query != query {
		return nil, fmt.Errorf("search results by query '%s' does'nt exists", query)
	}

	// updating expiration of video
	if err := s.expire(ctx, "search_result:"+query); err != nil {
		return nil, err
	}

	return searchResultDTO, nil
}

// put it in a cache
func (s *storage) CreateSearchResult(ctx context.Context, searchResult *SearchResult) error {

	if _, err := s.GetSearchResultByQuary(ctx, searchResult.Query); err == nil {
		return fmt.Errorf("search result bu queary '%s' already exists in cache", searchResult.Query)
	}

	// trying to write it into redis
	_, err := s.client.Pipelined(ctx, func(rdb cache.Pipeliner) error {

		// dto contains slice of video.id in place  of video models
		searchResultDTO := NewStoreSearchResultDTO(searchResult)
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

	// updating expiration of video
	if err := s.expire(ctx, "search_result:"+searchResult.Query); err != nil {
		return err
	}

	return err
}
