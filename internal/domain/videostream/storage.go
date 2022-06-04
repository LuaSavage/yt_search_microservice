package videostream

import (
	"context"
	"encoding/json"
	"fmt"

	cache "github.com/LuaSavage/yt_search_microservice/pkg/client/cache"
)

type Storage interface {
	Get(ctx context.Context, videoId string) (*VideoStreamPool, error)
	Create(ctx context.Context, streamPool *VideoStreamPool) error
}

type storage struct {
	client cache.Client
}

func NewStorage(client cache.Client) Storage {
	return &storage{client: client}
}

func (s *storage) Get(ctx context.Context, videoId string) (*VideoStreamPool, error) {
	// get search result hash by type:id from redis
	stringMap := s.client.HGetAll(ctx, "video_stream_pool:"+videoId)

	if stringMap.Err() != nil {
		return nil, stringMap.Err()
	}

	var streamPool VideoStreamPool = VideoStreamPool{
		VideoId: videoId,
		Streams: []VideoStream{},
	}

	if streamPool.VideoId != stringMap.Val()["videoId"] {
		return nil, fmt.Errorf("video stream pool by id '%s' does'nt exists", videoId)
	}

	// extracting video streams
	json.Unmarshal([]byte(stringMap.Val()["streams"]), &streamPool.Streams)

	return &streamPool, nil
}

// put it in a cache
func (s *storage) Create(ctx context.Context, streamPool *VideoStreamPool) error {

	if _, err := s.Get(ctx, streamPool.VideoId); err == nil {
		return fmt.Errorf("video stream pool by id '%s' already exists in cache", streamPool.VideoId)
	}

	// trying to write it into redis
	_, err := s.client.Pipelined(ctx, func(rdb cache.Pipeliner) error {

		cmd := rdb.HSet(ctx, "video_stream_pool:"+streamPool.VideoId, "videoId", streamPool.VideoId)

		if cmd.Err() != nil {
			return cmd.Err()
		}

		streamsMarshaled, _ := json.Marshal(streamPool.Streams)
		cmd = rdb.HSet(ctx, "video_stream_pool:"+streamPool.VideoId, "streams", streamsMarshaled)

		if cmd.Err() != nil {
			return cmd.Err()
		}

		return nil
	})

	return err
}
