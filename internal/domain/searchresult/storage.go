package searchresult

import (
	"context"
	"errors"
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
	return nil, errors.New(fmt.Sprintf("Error! Search results by query '%s' alrady exists", query))
}

// put it in a cache
func (s *storage) CreateSearchResult(ctx context.Context, searchResult *SearchResult) error {
	// someshit here

	_, err := s.GetSearchResultByQuary(ctx, searchResult.Query)

	if err == nil {
		return errors.New(fmt.Sprintf("Error! Search results by query '%s' alrady exists", searchResult.Query))
	}

	return errors.New(fmt.Sprintf("Error! Search results by query '%s' alrady exists", searchResult.Query))
}
