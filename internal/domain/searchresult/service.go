package searchresult

import (
	"context"
	"log"

	video "github.com/LuaSavage/yt_search_microservice/internal/domain/video"
	ytsearch "github.com/LuaSavage/yt_search_microservice/pkg/client/ytsearch"
)

type Service interface {
	GetSearchResultByQuary(ctx context.Context, query string) (*SearchResult, error)
	Search(ctx context.Context, query string) (*SearchResult, error)
}

type service struct {
	searchApi ytsearch.Service
	storage   Storage
}

func NewService(storage Storage, searchApi ytsearch.Service) Service {
	return &service{
		searchApi: searchApi,
		storage:   storage}
}

func (s *service) GetSearchResultByQuary(ctx context.Context, query string) (*SearchResult, error) {
	results, err := s.storage.GetSearchResultByQuary(ctx, query)

	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *service) Search(ctx context.Context, query string) (*SearchResult, error) {
	results, err := s.GetSearchResultByQuary(ctx, query)

	if err == nil {
		return results, nil
	}

	result, err := s.searchApi.Search(query)

	if err != nil {
		log.Fatal(err)
	}

	// Searching for new results from lib's api
	var videoPool []video.Video

	for _, res := range result {
		videoPool = append(videoPool, video.Video(res))
	}

	currentSearchResults := &SearchResult{Query: query, Videos: videoPool}

	return currentSearchResults, nil
}
