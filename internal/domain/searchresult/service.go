package searchresult

import (
	"context"
	"log"

	YtSearch "github.com/AnjanaMadu/YTSearch"
	video "github.com/LuaSavage/yt_search_microservice/internal/video"
)

type Service interface {
	GetSearchResultByQuary(ctx context.Context, query string) (*SearchResult, error)
	Search(ctx context.Context, query string) (*SearchResult, error)
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage}
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

	result, err := YtSearch.Search(query)

	if err != nil {
		log.Fatal(err)
	}

	// Searching for new results from lib's api
	var videoPool []video.Video

	for _, res := range result {

		video := video.Video{
			Id:          res.VideoId,
			Thumbnail:   res.Thumbnail,
			PublishTime: res.PublishTime,
			Channel:     res.Channel, Views: res.Views}

		videoPool = append(videoPool, video)
	}

	currentSearchResults := &SearchResult{Query: query, Videos: videoPool}

	return currentSearchResults, nil
}
