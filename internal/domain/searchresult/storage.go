package searchresult

import (
	"context"
	"fmt"
)

type Storage interface {
	GetSearchResultByQuary(ctx context.Context, query string) (*SearchResult, error)
	CreateSearchResult(ctx context.Context, searchResult *SearchResult) error
}

type storage struct{}

func NewStorage() Storage {
	return &storage{}
}

func (s *storage) GetSearchResultByQuary(ctx context.Context, query string) (*SearchResult, error) {
	return nil, fmt.Errorf("search results by query '%s' alrady exists", query)
}

// put it in a cache
func (s *storage) CreateSearchResult(ctx context.Context, searchResult *SearchResult) error {

	_, err := s.GetSearchResultByQuary(ctx, searchResult.Query)

	if err != nil {
		return err
	}

	// some redis shit here

	return nil
}
