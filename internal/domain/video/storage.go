package video

import (
	"context"
	"encoding/json"

	cache "github.com/LuaSavage/yt_search_microservice/pkg/client/cache"
	redis "github.com/go-redis/redis/v8"
)

type Storage interface {
	GetVideoByID(ctx context.Context, id string) (*Video, error)
	//SetVideo(ctx context.Context, video Video) error
	CreateVideo(ctx context.Context, video Video) error
}

type storage struct {
	client cache.Client
}

func NewStorage(client cache.Client) Storage {
	return &storage{client: client}
}

func (s *storage) GetVideoByID(ctx context.Context, id string) (*Video, error) {
	var video Video

	// get video hash by type:id from redis
	if err := s.client.HGetAll(ctx, "video:"+id).Scan(&video); err != nil {
		return nil, err
	}

	return &video, nil
}

func (s *storage) CreateVideo(ctx context.Context, video Video) error {

	if _, err := s.GetVideoByID(ctx, video.Id); err == nil {
		return err
	}

	// trying to write it into redis
	if _, err := s.client.Pipelined(ctx, func(rdb redis.Pipeliner) error {
		var videoMaped map[string]interface{}

		data, _ := json.Marshal(video)
		json.Unmarshal(data, &videoMaped)

		for key, value := range videoMaped {
			rdb.HSet(ctx, "video:"+video.Id, key, value)
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
