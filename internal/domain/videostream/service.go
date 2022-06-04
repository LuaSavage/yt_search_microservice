package videostream

import (
	"context"
	"fmt"

	ytvideo "github.com/LuaSavage/yt_search_microservice/pkg/client/ytvideo"
)

type Service interface {
	Get(ctx context.Context, videoId string) (*VideoStreamPool, error)
	Create(ctx context.Context, videoId string) error
}

type service struct {
	storage       Storage
	ytVideoClient ytvideo.Client
}

func NewService(storage Storage, ytVideoClient ytvideo.Client) Service {
	return &service{
		storage:       storage,
		ytVideoClient: ytVideoClient,
	}
}

func (s *service) Create(ctx context.Context, videoId string) error {
	_, err := s.Get(ctx, videoId)

	if err != nil {
		return err
	}

	//case of unexisting video in cache
	//we better request it by iteslf
	ytVideo, err := s.ytVideoClient.GetVideoContext(ctx, videoId)

	if err != nil {
		return err
	} else {
		streamPool := VideoStreamPool{
			VideoId: videoId,
			Streams: []VideoStream{},
		}

		for _, format := range ytVideo.Formats.Type("mp4") {
			streamPool.Streams = append(streamPool.Streams, VideoStream{
				VideoId: videoId,
				Quality: format.QualityLabel,
				Url:     format.URL,
			})
		}

		s.storage.Create(ctx, &streamPool)
	}

	return nil
}

func (s *service) Get(ctx context.Context, videoId string) (*VideoStreamPool, error) {
	video, err := s.storage.Get(ctx, videoId)

	if err != nil {
		return nil, fmt.Errorf("video stream pool by id '%s' does'nt exists", videoId)
	}

	return video, nil
}
