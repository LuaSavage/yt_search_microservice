package video

import (
	"context"
	"fmt"
)

type Service interface {
	GetVideoByID(ctx context.Context, id string) (*Video, error)
	CreateVideo(ctx context.Context, video Video) error
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage}
}

func (s *service) CreateVideo(ctx context.Context, video Video) error {
	return s.storage.CreateVideo(ctx, video)
}

func (s *service) GetVideoByID(ctx context.Context, id string) (*Video, error) {
	video, err := s.storage.GetVideoByID(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("video by id '%s' does'nt exists", id)
	}

	return video, nil
}
