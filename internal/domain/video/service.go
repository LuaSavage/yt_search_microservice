package video

import "context"

type Service interface {
	GetVideoByID(ctx context.Context, id string)
	CreateVideo(ctx context.Context, video Video)
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage}
}

func (s *service) CreateVideo(ctx context.Context, video Video)
func (s *service) GetVideoByID(ctx context.Context, id string) {
	panic("unimplemented")
}
