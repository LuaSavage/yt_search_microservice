package videostream

import (
	"context"
	"fmt"
)

type Service interface {
	Get(ctx context.Context, videoId string) (*VideoStreamPool, error)
	Create(ctx context.Context, video *VideoStreamPool) error
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage}
}

func (s *service) Create(ctx context.Context, video *VideoStreamPool) error {
	return s.storage.Create(ctx, video)
}

func (s *service) Get(ctx context.Context, videoId string) (*VideoStreamPool, error) {
	video, err := s.storage.Get(ctx, videoId)

	if err != nil {
		return nil, fmt.Errorf("video stream pool by id '%s' does'nt exists", videoId)
	}

	return video, nil
}
