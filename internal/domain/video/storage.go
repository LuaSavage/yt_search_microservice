package video

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	cache "github.com/LuaSavage/yt_search_microservice/pkg/client/cache"
)

type Storage interface {
	GetVideoByID(ctx context.Context, id string) (*Video, error)
	//SetVideo(ctx context.Context, video Video) error
	CreateVideo(ctx context.Context, video *Video) error
}

type storage struct {
	client     cache.Client
	expiration int
}

func NewStorage(client cache.Client, expiration int) Storage {
	return &storage{
		client:     client,
		expiration: expiration,
	}
}

func (s *storage) expire(ctx context.Context, key string) error {
	if cmd := s.client.Expire(ctx, key, time.Second*time.Duration(s.expiration)); cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}

func (s *storage) GetVideoByID(ctx context.Context, id string) (*Video, error) {
	var video *Video = &Video{}

	// get video hash by type:id from redis
	if err := s.client.HGetAll(ctx, "video:"+id).Scan(video); err != nil {
		return nil, err
	}

	if video.Id != id {
		return nil, fmt.Errorf("retriven video holds inproper id or none. video id: %s", video.Id)
	}

	// updating expiration of video
	if err := s.expire(ctx, "video:"+id); err != nil {
		return nil, err
	}

	return video, nil
}

func (s *storage) CreateVideo(ctx context.Context, video *Video) error {

	if _, err := s.GetVideoByID(ctx, video.Id); err == nil {
		return fmt.Errorf("video with id '%s' already exists in cache", video.Id)
	}

	var videoMaped map[string]interface{}
	data, _ := json.Marshal(video)
	json.Unmarshal(data, &videoMaped)

	// trying to write it into redis
	if _, err := s.client.Pipelined(ctx, func(rdb cache.Pipeliner) error {

		for key, value := range videoMaped {
			cmd := rdb.HSet(ctx, "video:"+video.Id, key, value)
			if cmd.Err() != nil {
				return cmd.Err()
			}
		}

		return nil

	}); err != nil {
		return err
	}

	// updating expiration of video
	if err := s.expire(ctx, "video:"+video.Id); err != nil {
		return err
	}

	return nil
}
