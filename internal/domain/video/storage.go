package video

import "context"

type Storage interface {
	GetVideoByID(ctx context.Context, id string) *Video
	SetVideo(ctx context.Context, video Video) error
	CreateVideo(ctx context.Context, video Video) error
}

type storage struct{}

func NewStorage() Storage {
	return &storage{}
}

func (s *storage) GetVideoByID(ctx context.Context, id string) *Video
func (s *storage) SetVideo(ctx context.Context, video Video) error
func (s *storage) CreateVideo(ctx context.Context, video Video) error
