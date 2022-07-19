package video

import (
	"context"
	"encoding/json"
	"fmt"
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

// initialise test object
var (
	testVideo *Video = &Video{
		Title:       "Test video",
		Id:          "=TestvdIeo",
		PublishTime: "22-09-2011",
		Channel:     "someshit",
		Views:       "228k",
		Thumbnail:   "htttp://sdf.dickpic",
	}
	ctx        context.Context = context.TODO()
	expiration int             = 100
)

func TestGetVideoByIDError(t *testing.T) {
	t.Run("error due to empty video object response", func(t *testing.T) {
		redisClient, _ := newTestRedis()
		videoStorage := NewStorage(redisClient, expiration)
		retrivedVideo, err := videoStorage.GetVideoByID(ctx, testVideo.Id)
		assert.Error(t, err, fmt.Sprintf("and retrived video looks like %+v\n", retrivedVideo))
	})
}

func TestGetVideoByIDOK(t *testing.T) {
	t.Run("error due to test video doesnt received", func(t *testing.T) {
		redisClient, redisMock := newTestRedis()
		videoStorage := NewStorage(redisClient, expiration)

		// writing test video into
		var videoMaped map[string]string

		testVideoMarshaled, _ := json.Marshal(testVideo)
		json.Unmarshal(testVideoMarshaled, &videoMaped)

		for key, value := range videoMaped {
			redisMock.HSet(ctx, "video:"+testVideo.Id, key, value)
		}

		retrivenVideo, err := videoStorage.GetVideoByID(ctx, testVideo.Id)
		assert.Equal(t, testVideo, *retrivenVideo, err)
	})
}

func TestCreateVideoOK(t *testing.T) {
	t.Run("error due to the video  dosn't properly created", func(t *testing.T) {
		redisClient, _ := newTestRedis()
		videoStorage := NewStorage(redisClient, expiration)

		_, err := videoStorage.GetVideoByID(ctx, testVideo.Id)
		assert.Error(t, err)

		err = videoStorage.CreateVideo(ctx, testVideo)
		assert.NoError(t, err)

		// If it succesfully written then it has to be same as origin
		retrivenVideo, err := videoStorage.GetVideoByID(ctx, testVideo.Id)
		assert.Equal(t, testVideo, *retrivenVideo, err)
	})
}

// prevention of double write issue
func TestCreateVideoError(t *testing.T) {
	t.Run("error due to double storring video up issue", func(t *testing.T) {
		redisClient, _ := newTestRedis()
		videoStorage := NewStorage(redisClient, expiration)

		err := videoStorage.CreateVideo(ctx, testVideo)
		assert.NoError(t, err)
		err = videoStorage.CreateVideo(ctx, testVideo)
		assert.Error(t, err)
	})
}
