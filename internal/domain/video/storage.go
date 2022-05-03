package video

import "context"

type Storage interface {
	GetVideoByID(ctx context.Context, id string) (*Video, error)
	//SetVideo(ctx context.Context, video Video) error
	CreateVideo(ctx context.Context, video Video) error
}

type storage struct{}

func NewStorage() Storage {
	return &storage{}
}

func (s *storage) GetVideoByID(ctx context.Context, id string) (*Video, error) {
	// get from redis
	video := Video{}
	return &video, nil
}

func (s *storage) CreateVideo(ctx context.Context, video Video) error {
	_, err := s.GetVideoByID(ctx, video.Id)

	if err != nil {
		return err
	}

	// some redis shit

	return nil
}
