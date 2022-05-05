package video

import (
	"context"
	"encoding/json"
	"testing"

	miniredis "github.com/alicebob/miniredis"
	redismock "github.com/elliotchance/redismock/v8"
	redis "github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

// newTestRedis returns a redis.Cmdable.
func newTestRedis() (*redis.Client, *redismock.ClientMock) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return client, redismock.NewNiceMock(client)
}

func TestGetVideoByID(t *testing.T) {
	t.Run("error due to test video doesnt received", func(t *testing.T) {
		redisClient, redisMock := newTestRedis()
		videoStorage := NewStorage(redisClient)
		ctx := context.TODO()

		// initialise test object
		testVideo := Video{
			Title:       "Test video",
			Id:          "=TestvdIeo",
			PublishTime: "22-09-2011",
			Channel:     "someshit",
			Views:       "228k",
			Thumbnail:   "htttp://sdf.dickpic",
		}

		// writing test video into
		var videoMaped map[string]string

		testVideoMarshaled, _ := json.Marshal(testVideo)
		json.Unmarshal(testVideoMarshaled, &videoMaped)

		for key, value := range videoMaped {
			redisMock.HSet(ctx, "video:"+testVideo.Id, key, value)
		}

		retrivenVideo, err := videoStorage.GetVideoByID(context.Background(), testVideo.Id)
		assert.Equal(t, testVideo, *retrivenVideo, err)
	})
}
